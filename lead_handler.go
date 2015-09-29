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
	name := c.Form("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "name is mandatory")
	}
	hashCode := c.Form("hashCode")
	return h.ds.Create(name, hashCode)
}

// CountByInvites count leads by invite
func (h LeadHandler) CountByInvites(c *echo.Context) error {
	hashCode := c.Form("hashCode")
	if hashCode == "" {
		return c.String(http.StatusBadRequest, "hashCode is mandatory")
	}
	result, err := h.ds.CountByInvites(hashCode)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
