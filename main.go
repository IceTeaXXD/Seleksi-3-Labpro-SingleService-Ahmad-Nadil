package main

import (
	"singleservice/controllers"
	"singleservice/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://chatakudong.vercel.app","http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// POST
	r.POST("/login", controllers.Login)
	// r.POST("/barang", controllers.CreateBarang)
	// r.post("/perusahaan", controllers.CreatePerusahaan)

	// // GET
	// r.GET("/self", controllers.GetSelf)
	// r.GET("/barang", controllers.GetBarang)
	// r.GET("/barang/:id", controllers.GetBarangByID)
	// r.get("/perusahaan", controllers.GetPerusahaan)
	// r.get("/perusahaan/:id", controllers.GetPerusahaanByID)

	// // DELETE
	// r.DELETE("/barang/:id", controllers.DeleteBarang)
	// r.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)

	// // PUT
	// r.PUT("/barang/:id", controllers.UpdateBarang)
	// r.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)

	r.Run()
}