package Business

import (
	"net/http"
	"fmt"
	"log"
	"TestProject/Models"
	"TestProject/DataAccess"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func Ping(w http.ResponseWriter, r *http.Request){
	ResponseWithJSON(w, nil, http.StatusOK)
}

func GetTeams(w http.ResponseWriter, r *http.Request){
	teams,err := DataAccess.GetTeams()
	if err != nil {
		ErrorWithJSON(w, "DataAccess error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.MarshalIndent(teams, "", "  ")
	if err != nil {
		ErrorWithJSON(w, "Marshal error", http.StatusInternalServerError)
		log.Println("Marshal error: ", err)
		return
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func GetTeam(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	var team *Models.Team
	team,err := DataAccess.GetTeam(id)
	if err != nil {
		ErrorWithJSON(w, "DataAccess error", http.StatusInternalServerError)
		return
	}

	if team.Id == "" {
		ErrorWithJSON(w, "Team not found", http.StatusNotFound)
		return
	}

	respBody, err := json.MarshalIndent(team, "", "  ")
	if err != nil {
		ErrorWithJSON(w, "Marshal error", http.StatusInternalServerError)
		log.Println("Marshal error: ", err)
		return
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team Models.Team
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&team)
	if err != nil {
		ErrorWithJSON(w, err.Error(), http.StatusBadRequest)
		log.Println("Decode error: ", err)
		return
	}

	err = DataAccess.CreateTeam(team)
	if err != nil {
		ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
		log.Println("Create error: ", err)
		return
	}

	ResponseWithJSON(w, nil, http.StatusOK)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var team Models.Team
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&team)
	if err != nil {
		ErrorWithJSON(w, err.Error(), http.StatusBadRequest)
		log.Println("Decode error: ", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	err = DataAccess.UpdateTeam(id, team)
	if err != nil {
		ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
		log.Println("Update error: ", err)
		return
	}

	ResponseWithJSON(w, nil, http.StatusOK)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	err := DataAccess.DeleteTeam(id)
	if err != nil {
		ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
		log.Println("Delete error: ", err)
		return
	}

	ResponseWithJSON(w, nil, http.StatusOK)
}

func Draw(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	minCountStr := vars["minCount"]
	minCount, err := strconv.Atoi(minCountStr)
	if err != nil {
		ErrorWithJSON(w, "Input parameter error", http.StatusInternalServerError)
		return
	}

	teams,err := DataAccess.GetTeams()
	if err != nil {
		ErrorWithJSON(w, "DataAccess error", http.StatusInternalServerError)
		return
	}

	groups := GetGroups(teams, minCount)

	respBody, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		ErrorWithJSON(w, "Marshal error", http.StatusInternalServerError)
		log.Println("Marshal error: ", err)
		return
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func CloseSession(){
	DataAccess.CloseSession()
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{\"message\": %q}", message)
}
