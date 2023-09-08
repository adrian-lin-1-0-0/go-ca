/*
 * Swagger Example
 *
 * Swagger Example
 *
 * API version: 1.0.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	Route{
		"AddPet",
		strings.ToUpper("Post"),
		"/v2/pet",
		AddPet,
	},

	Route{
		"DeletePet",
		strings.ToUpper("Delete"),
		"/v2/pet/{petId}",
		DeletePet,
	},

	Route{
		"GetPetById",
		strings.ToUpper("Get"),
		"/v2/pet/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		strings.ToUpper("Put"),
		"/v2/pet",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		strings.ToUpper("Post"),
		"/v2/pet/{petId}",
		UpdatePetWithForm,
	},
}
