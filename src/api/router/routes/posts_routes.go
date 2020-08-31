package routes

import (
	"../../controllers"
	"net/http"
)

var postsRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
		AuthRequired:false,
	},

	Route{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
		AuthRequired:true,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetPost,
		AuthRequired:true,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
		AuthRequired:false,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletPost,
		AuthRequired:true,
	},
}