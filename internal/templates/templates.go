package templates

import (
	"html/template"
)

func BundleBoilerplate(path string) *template.Template {
	return template.Must(template.ParseFiles(
		path,
		"internal/templates/header.html.tmpl",
		"internal/templates/footer.html.tmpl",
	))
}
