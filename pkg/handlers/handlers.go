package handlers

import (
	"net/http"

	"github.com/dh-n/go-course/pkg/config"
	"github.com/dh-n/go-course/pkg/models"
	"github.com/dh-n/go-course/pkg/render"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "Hello world")

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there from template data - Yeehaw (Top Gun)"
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "Hello world")
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}


