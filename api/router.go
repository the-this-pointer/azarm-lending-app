package api

import (
	"azarm-lending-backend/api/middleware"
	ApiRouter "azarm-lending-backend/api/router"
	"azarm-lending-backend/api/user"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	//append user routes
	ApiRouter.AppRoutes = append(ApiRouter.AppRoutes, user.Routes)

	for _, route := range ApiRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//check to see if route should be protected with jwt
			if r.Protected {
				handler = middleware.JWTMiddleware(r.HandlerFunc)
			}

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}

	}

	return router
}
