package Route

import "net/http"
import "TestProject/Business"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Ping",
		"GET",
		"/",
		Business.Ping,
	},
	Route{
		"GetTeams",
		"GET",
		"/teams",
		Business.GetTeams,
	},
	Route{
		"CreateTeam",
		"POST",
		"/teams",
		Business.CreateTeam,
	},
	Route{
		"GetTeam",
		"GET",
		"/teams/{id}",
		Business.GetTeam,
	},
	Route{
		"UpdateTeam",
		"PUT",
		"/teams/{id}",
		Business.UpdateTeam,
	},
	Route{
		"DeleteTeam",
		"DELETE",
		"/teams/{id}",
		Business.DeleteTeam,
	},
	Route{
		"Draw",
		"GET",
		"/draw/{minCount}",
		Business.Draw,
	},
}
