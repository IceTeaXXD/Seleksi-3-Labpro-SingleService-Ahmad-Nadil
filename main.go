package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"singleservice/auth"
	"singleservice/controllers"
	"singleservice/initializers"
	"singleservice/migrate"
	// "fmt"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// run the migration
	migrate.Migrate()

	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins: []string{"http://localhost:3000", "https://ohl-fe.vercel.app", "http://localhost:5173", "https://monolith-labpro.up.railway.app/", "http://127.0.0.1:8000/"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "Content-Length", "X-CSRF-Token", "Authorization"},
	}))

	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("/self", controllers.GetSelf)
		authorized.GET("/perusahaan", controllers.GetPerusahaan)
		authorized.GET("/perusahaan/:id", controllers.GetPerusahaanByID)
		authorized.POST("/barang", controllers.CreateBarang)
		authorized.POST("/perusahaan", controllers.CreatePerusahaan)
		authorized.DELETE("/barang/:id", controllers.DeleteBarang)
		authorized.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)
		authorized.PUT("/barang/:id", controllers.UpdateBarang)
		authorized.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)
	}

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/buy", controllers.BuyBarang);
	r.GET("/barang", controllers.GetBarang)
	r.GET("/barang/:id", controllers.GetBarangByID)
	r.OPTIONS("/login", controllers.LoginOptions)

	r.Run()
}
