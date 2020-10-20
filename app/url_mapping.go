package app

import(
	"github.com/arya0618/my-bookstore-api/controllers/ping"
	"github.com/arya0618/my-bookstore-api/controllers/users"
)

func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.POST("/users",users.CreatetUser)
	router.GET("/users/:user_id",users.GetUser)
//	router.GET("/users/search",controllers.SearchUser)
}