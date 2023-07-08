package handlers

import (
	"github.com/kabilovtoha/go_study_bookings/pkg/config"
	"github.com/kabilovtoha/go_study_bookings/pkg/models"
	"github.com/kabilovtoha/go_study_bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) (r *Repository) {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the Home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "index.page.gohtml", &models.TemplateData{})
}

// /rooms/generals-quarters
// /rooms/majors-suite

// About is the About page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	StringMap := map[string]string{"remote_ip": remoteIP}

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.gohtml", &models.TemplateData{})
}

func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

func (repo *Repository) Contacts(w http.ResponseWriter, r *http.Request) {
	StringMap := map[string]string{}
	render.RenderTemplate(w, "contacts.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) RoomsGeneralsQuarters(w http.ResponseWriter, r *http.Request) {

	StringMap := map[string]string{
		"post_header": "General's Quarters",
		"image":       "/static/images/room-images/generals-quarters.png",
	}

	render.RenderTemplate(w, "rooms.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) RoomsMajorsSuite(w http.ResponseWriter, r *http.Request) {
	StringMap := map[string]string{
		"post_header": "Major's Quarters",
		"image":       "/static/images/room-images/marjors-suite.png",
	}
	render.RenderTemplate(w, "rooms.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}
