package intented

import (
	"net/http"

	"github.com/Xmio/intented/datastores"
	"github.com/labstack/echo"
)

// AppHandler contains all /app handlers
type AppHandler struct {
	ds datastores.App
}

// NewAppHandler creates a echo handler for /app
func NewAppHandler(ds datastores.App) AppHandler {
	return AppHandler{ds}
}

// Create a new app in database
func (h AppHandler) Create(c *echo.Context) error {
	name := c.Form("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "name is mandatory")
	}
	return h.ds.Create(name)
}
