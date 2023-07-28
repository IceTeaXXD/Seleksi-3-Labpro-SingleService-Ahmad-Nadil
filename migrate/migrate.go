package migrate

import (
	"singleservice/initializers"
	model "singleservice/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate(){
	fmt.Println("Migrating........................")
	if !initializers.DB.Migrator().HasTable(&model.User{}) {
		fmt.Println("Migrating User........................")
		initializers.DB.AutoMigrate(&model.User{})
		initializers.DB.Create(&model.User{ Username: "admin", Password: "admin", Name: "Administrator", IsAdmin: true })
	}
	if !initializers.DB.Migrator().HasTable(&model.Perusahaan{}) {
		fmt.Println("Migrating Perusahaan........................")
		initializers.DB.AutoMigrate(&model.Perusahaan{})
		initializers.DB.Create(&model.Perusahaan{ Nama: "PT. Maju Mundur", Alamat: "Jl. Maju Mundur No. 1", NoTelepon: "08123456789", KodePajak: "MU1" })
	}
	if !initializers.DB.Migrator().HasTable(&model.Barang{}) {
		fmt.Println("Migrating Barang........................")
		initializers.DB.AutoMigrate(&model.Barang{})
		initializers.DB.Create(&model.Barang{ Nama: "Beras", Harga: 10000, Stok: 100, PerusahaanID: "1", KodeBarang: "BR002" })
	}
}
