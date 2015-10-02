package intented

import (
	"net/http"
	"time"

	"github.com/Xmio/intented/datastores"
	"github.com/labstack/echo"
)

// LeadHandler contains all /lead handlers
type LeadHandler struct {
	ds datastores.Lead
}

// NewLeadHandler creates a echo handler for /lead
func NewLeadHandler(ds datastores.Lead) LeadHandler {
	return LeadHandler{ds}
}

// Create a new lead in database
func (h LeadHandler) Create(c *echo.Context) error {
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
func (h LeadHandler) CountByInvites(c *echo.Context) error {
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
