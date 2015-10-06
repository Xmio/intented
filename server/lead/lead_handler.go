package lead

import (
	"net/http"
	"time"

	"github.com/Xmio/intented/server/datastores"
	"github.com/labstack/echo"
)

// Handler contains all /lead handlers
type Handler struct {
	ds datastores.Lead
}

// NewHandler creates a echo handler for /lead
func NewHandler(ds datastores.Lead) Handler {
	return Handler{ds}
}

// Create a new lead in database
func (h Handler) Create(c *echo.Context) error {
	mail := c.Form("mail")
	if mail == "" {
		return c.String(http.StatusBadRequest, "mail is mandatory")
	}

	result, err := h.ds.GetHashByMail(mail)
	if result != "" {
		return c.JSON(http.StatusOK, result)
	}

	invited := c.Form("invited")
	h.ds.Create(mail, invited)

	result, err = h.ds.GetHashByMail(mail)
	if err != nil {
		return err
	}

	cookie := http.Cookie{Name: "ownHash", Value: result, Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(c.Response().Writer(), &cookie)

	return c.JSON(http.StatusOK, result)
}

// CountByInvites count leads by invite
func (h Handler) CountByInvites(c *echo.Context) error {
	hashCode := c.Param("hashCode")
	if hashCode == "" {
		return c.String(http.StatusBadRequest, "hashCode is mandatory")
	}
	count, err := h.ds.CountByInvites(hashCode)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, count)
}
