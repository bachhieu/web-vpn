package main

import (
	"bachhieu/web-vpn/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	
	router := httprouter.New()
	
	router.GET("/get-vpn", routes.GetAllVpns)
	// router.POST("/auth/login", routes.Login)
	// router.POST("/auth/register", routes.Register)

	// router.GET("/posts", middlewares.CheckJwt(routes.GetAllPosts))
	// router.GET("/me/posts", middlewares.CheckJwt(routes.GetMyPosts))
	// router.POST("/posts", middlewares.CheckJwt(routes.CreatePost))
	// router.PUT("/posts/:id", middlewares.CheckJwt(routes.EditPost))
	// router.DELETE("/posts/:id", middlewares.CheckJwt(routes.DeletePost))

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
