package application

import (
	"github.com/waytkheming/bookstore-users-api/controllers"
)

func mapUrls() {

	router.GET("/users/:user_id", controllers.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)

}
