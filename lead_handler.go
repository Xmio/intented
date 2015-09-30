package intented

import (
	"net/http"

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
	hashCode := c.Form("hashCode")
	return h.ds.Create(mail, hashCode)
}

// CountByInvites count leads by invite
func (h LeadHandler) CountByInvites(c *echo.Context) error {
	hashCode := c.Param("hashCode")
	if hashCode == "" {
		return c.String(http.StatusBadRequest, "hashCode is mandatory")
	}
	result, err := h.ds.CountByInvites(hashCode)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
