package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/kabilovtoha/go_study_bookings/internal/config"
	"github.com/kabilovtoha/go_study_bookings/internal/models"
	"github.com/kabilovtoha/go_study_bookings/internal/render"
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

	render.RenderTemplate(w, r, "index.page.gohtml", &models.TemplateData{})
}

// /rooms/generals-quarters
// /rooms/majors-suite

// About is the About page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	StringMap := map[string]string{"remote_ip": remoteIP}

	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
}
func (repo *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//render.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
	fmt.Println(r)
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	w.Write([]byte(fmt.Sprintf("PostAvailability, start_date is %s and end_date is %s", startDate, endDate)))
}

type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (repo *Repository) ApiPostAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	fmt.Println("startDate: ", startDate)
	fmt.Println("endDate: ", endDate)
	resp := JsonResponse{
		OK:      true,
		Message: "Available",
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-search-availability.page.gohtml", &models.TemplateData{})
}

func (repo *Repository) Contacts(w http.ResponseWriter, r *http.Request) {
	StringMap := map[string]string{}
	render.RenderTemplate(w, r, "contacts.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) RoomsGeneralsQuarters(w http.ResponseWriter, r *http.Request) {

	StringMap := map[string]string{
		"post_header": "General's Quarters",
		"image":       "/static/images/room-images/generals-quarters.png",
		"room_type":   "gq",
	}

	render.RenderTemplate(w, r, "rooms.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}

func (repo *Repository) RoomsMajorsSuite(w http.ResponseWriter, r *http.Request) {
	StringMap := map[string]string{
		"post_header": "Major's Quarters",
		"image":       "/static/images/room-images/marjors-suite.png",
		"room_type":   "mq",
	}
	render.RenderTemplate(w, r, "rooms.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})
}
