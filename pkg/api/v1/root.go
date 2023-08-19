package v1

import (
	"html/template"
	"net/http"

	"github.com/john-odonnell/psite/v2/internal/templates"
)

// RootHandler implements the http.Handler interface to serve the rendered
// template of the server's root.
type RootHandler struct {
	template *template.Template
}

// DefaultRootHandler returns a new RootHandler instance with the default
// template for the server's root.
func DefaultRootHandler() RootHandler {
	path := "internal/templates/index.html.tmpl"
	return RootHandler{
		template: templates.BundleBoilerplate(path),
	}
}

// ServeHTTP renders the RootHandler's template to the provided ResponseWriter.
func (h RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.template.Execute(w, nil)
}
