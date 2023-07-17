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
		AllowOrigins: []string{"http://localhost:3000", "https://ohl-fe.vercel.app","http://localhost:5173","https://monolith-production-63a6.up.railway.app"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "Content-Length", "X-CSRF-Token", "Authorization"},
	}))

	// POST
	r.POST("/login", controllers.Login)
	r.POST("/barang", controllers.CreateBarang)
	r.POST("/perusahaan", controllers.CreatePerusahaan)
	r.POST("/register", controllers.Register)
	r.POST("/transaksi", controllers.CreateTransaksi)
	
	// GET
	r.GET("/self", controllers.GetSelf)
	r.GET("/barang", controllers.GetBarang)
	r.GET("/barang/:id", controllers.GetBarangByID)
	r.GET("/perusahaan", controllers.GetPerusahaan)
	r.GET("/perusahaan/:id", controllers.GetPerusahaanByID)
	r.GET("/transaksi/:username", controllers.GetTransaksiByUser)
	
	// DELETE
	r.DELETE("/barang/:id", controllers.DeleteBarang)
	r.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)
	
	// PUT
	r.PUT("/barang/:id", controllers.UpdateBarang)
	r.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)

	// OPTIONS
	r.OPTIONS("/login", controllers.LoginOptions)

	r.Run()
}