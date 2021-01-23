package routes

import "net/http"
import "../../controllers"

var postsRoutes = []Route{
	{
		Uri:    "/posts",
		Method: http.MethodGet,
		Handler: controllers.GetPosts,
	},
	{
		Uri:    "/posts",
		Method: http.MethodPost,
		Handler: controllers.CreatePost,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetPost,
	},
	{
		Uri:    "/posts/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdatePost,
	},
	{
		Uri:    "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeletePost,
	},
}
