package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) unbind(w http.ResponseWriter, r *http.Request) {
	instanceID := mux.Vars(r)["instance_id"]
	log.Println(instanceID)
	bindingID := mux.Vars(r)["biding_id"]
	log.Println(bindingID)
	w.Write(responseEmptyJSON)
	// TODO: Returns 200 or 410; also see spec for response body format
	w.WriteHeader(http.StatusOK)
}