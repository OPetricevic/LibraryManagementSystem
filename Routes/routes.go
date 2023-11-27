package routes

import (
	"net/http"

	handlers "github.com/OPetricevic/LibraryManagementSystem/Handlers"
)

func RegisterRoutes(mux *http.ServeMux, uc *handlers.UserController) {
	mux.Handler("register", uc.Register)
}
