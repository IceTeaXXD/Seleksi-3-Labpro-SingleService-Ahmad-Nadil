package main

import (
	"singleservice/controllers"
	"singleservice/initializers"
	"singleservice/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://ohl-fe.vercel.app","http://localhost:5173","https://monolith-production-63a6.up.railway.app"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "Content-Length", "X-CSRF-Token", "Authorization"},
	}))

	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("/self", controllers.GetSelf)
		authorized.GET("/barang", controllers.GetBarang)
		authorized.GET("/barang/:id", controllers.GetBarangByID)
		authorized.GET("/perusahaan", controllers.GetPerusahaan)
		authorized.GET("/perusahaan/:id", controllers.GetPerusahaanByID)
		authorized.GET("/transaksi/:username", controllers.GetTransaksiByUser)
		authorized.POST("/barang", controllers.CreateBarang)
		authorized.POST("/perusahaan", controllers.CreatePerusahaan)
		authorized.POST("/transaksi", controllers.CreateTransaksi)
		authorized.DELETE("/barang/:id", controllers.DeleteBarang)
		authorized.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)
		authorized.PUT("/barang/:id", controllers.UpdateBarang)
		authorized.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)
	}

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.OPTIONS("/login", controllers.LoginOptions)

	r.Run()
}