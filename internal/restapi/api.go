package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/fantasyFootballDraftGo/internal/db/migrate/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type RESTAPI struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *RESTAPI {
	return &RESTAPI{
		DB: db,
	}
}

func (api *RESTAPI) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *RESTAPI) createPick(w http.ResponseWriter, r *http.Request) {
	var pick *models.Pick
	err := json.NewDecoder(r.Body).Decode(&pick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	api.DB.Create(&pick)
	w.WriteHeader(http.StatusAccepted)
}

func (api *RESTAPI) getPicks(w http.ResponseWriter, r *http.Request) {
	var picks []*models.Pick
	query := r.URL.Query()
	manager := query.Get("manager")
	api.DB.Where("manager=?", manager).Find(&picks)
	json.NewEncoder(w).Encode(&picks)
}

func (api *RESTAPI) HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.health).Methods("GET")
	r.HandleFunc("/pick", api.createPick).Methods("POST")
	r.HandleFunc("/results", api.getPicks).Methods("GET")
	http.ListenAndServe(":8080", r)
}
