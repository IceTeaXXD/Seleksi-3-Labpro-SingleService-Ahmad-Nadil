package controllers

import (
	"singleservice/initializers"
	model "singleservice/models"
	// "singleservice/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSelf(c *gin.Context) {
	// get current user from context
	// user, _ := c.Get("user")
	// if !exists {
	//     c.JSON(http.StatusUnauthorized, gin.H{
	//         "status":  "error",
	//         "message": "User not authenticated",
	//         "data":    nil,
	//     })
	//     return
	// }

	// return user data
	// c.JSON(http.StatusOK, gin.H{
	//     "status":  "success",
	//     "message": "User data retrieved successfully",
	//     "data": gin.H{
	//         "username": user.(*model.User).Username,
	//         "name":     user.(*model.User).Name,
	//     },
	// })
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User data retrieved successfully",
		"data": gin.H{
			"username": "admin",
			"name":     "administrator",
		},
	})
}

func GetBarang(c *gin.Context) {
	// get query parameters
	q := c.Query("q")
	perusahaan := c.Query("perusahaan")

	// retrieve barang from database
	var barangs []model.Barang
	query := initializers.DB
	if q != "" {
		query = query.Where("nama LIKE ? OR kode LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	if perusahaan != "" {
		query = query.Where("perusahaan_id = ?", perusahaan)
	}
	result := query.Find(&barangs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve barang",
			"data":    nil,
		})
		return
	}

	// return barang data
	var data []gin.H
	for _, barang := range barangs {
		data = append(data, gin.H{
			"id":            barang.ID,
			"nama":          barang.Nama,
			"harga":         barang.Harga,
			"stok":          barang.Stok,
			"kode":          barang.KodeBarang,
			"perusahaan_id": barang.PerusahaanID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang retrieved successfully",
		"data":    data,
	})
}

func GetBarangByID(c *gin.Context) {
	// retrieve barang ID from request
	barangID := c.Param("id")

	// retrieve barang from database
	var barang model.Barang
	result := initializers.DB.First(&barang, barangID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve barang",
			"data":    nil,
		})
		return
	}

	// return barang data
	data := gin.H{
		"id":            barang.ID,
		"nama":          barang.Nama,
		"harga":         barang.Harga,
		"stok":          barang.Stok,
		"kode":          barang.KodeBarang,
		"perusahaan_id": barang.PerusahaanID,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang retrieved successfully",
		"data":    data,
	})
}

func GetPerusahaan(c *gin.Context) {
	// get query parameters
	q := c.Query("q")

	// retrieve perusahaan from database
	var perusahaans []model.Perusahaan
	query := initializers.DB
	if q != "" {
		query = query.Where("nama LIKE ? OR kode_pajak LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	result := query.Find(&perusahaans)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve perusahaan",
			"data":    nil,
		})
		return
	}

	// return perusahaan data
	var data []gin.H
	for _, perusahaan := range perusahaans {
		data = append(data, gin.H{
			"id":      perusahaan.ID,
			"nama":    perusahaan.Nama,
			"alamat":  perusahaan.Alamat,
			"no_telp": perusahaan.NoTelepon,
			"kode":    perusahaan.KodePajak,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Perusahaan retrieved successfully",
		"data":    data,
	})
}

func GetPerusahaanByID(c *gin.Context) {
	// retrieve perusahaan ID from request
	perusahaanID := c.Param("id")

	// retrieve perusahaan from database
	var perusahaan model.Perusahaan
	result := initializers.DB.First(&perusahaan, perusahaanID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve perusahaan",
			"data":    nil,
		})
		return
	}

	// return perusahaan data
	data := gin.H{
		"id":      perusahaan.ID,
		"nama":    perusahaan.Nama,
		"alamat":  perusahaan.Alamat,
		"no_telp": perusahaan.NoTelepon,
		"kode":    perusahaan.KodePajak,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Perusahaan retrieved successfully",
		"data":    data,
	})
}

func GetTransaksiByUser(c *gin.Context) {
	// get id_pembeli from request parameters
	user_pembeli := c.Param("username")
	fmt.Println(user_pembeli)

	var transactions []model.Transaksi
	result := initializers.DB.Where("user_pembeli = ?", user_pembeli).Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve transactions",
			"data":    nil,
		})
		return
	}

	// return transaksi data
	var data []gin.H
	for _, transaction := range transactions {
		data = append(data, gin.H{
			"id":            transaction.ID,
			"user_pembeli":  transaction.UserPembeli,
			"nama_barang":   transaction.NamaBarang,
			"jumlah_barang": transaction.JumlahBarang,
			"total_harga":   transaction.TotalHarga,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Transactions retrieved successfully",
		"data":    data,
	})
}
