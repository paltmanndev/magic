package handlers

import (
	"net/http"

	"github.com/paltmanndev/magic/pkg/render"
)

// Home ist der Homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About ist der Aboutpage handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
