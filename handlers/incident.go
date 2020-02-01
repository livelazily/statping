package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func apiAllIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	incidents := core.AllIncidents()
	returnJson(incidents, w, r)
}

func apiCreateIncidentHandler(w http.ResponseWriter, r *http.Request) {
	var incident *types.Incident
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incident)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	newIncident := core.ReturnIncident(incident)
	_, err = newIncident.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(newIncident, "create", w, r)
}

func apiIncidentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incident, err := core.SelectIncident(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&incident)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	_, err = incident.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "update", w, r)
}

func apiDeleteIncidentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incident, err := core.SelectIncident(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = incident.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "delete", w, r)
}
