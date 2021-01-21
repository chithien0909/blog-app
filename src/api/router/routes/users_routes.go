package routes

import "net/http"
import "../../controllers"

var usersRoutes = []Route{
	{
		Uri:    "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		Uri:    "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}
