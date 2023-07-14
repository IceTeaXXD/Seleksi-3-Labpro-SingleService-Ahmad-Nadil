package controllers

import (
	"singleservice/initializers"
	model "singleservice/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func UpdateBarang(c *gin.Context) {
    // retrieve barang ID from request
    barangID := c.Param("id")

    // retrieve existing barang from database
    var existingBarang model.Barang
    result := initializers.DB.First(&existingBarang, barangID)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to retrieve existing barang",
            "data":    nil,
        })
        return
    }

    // parse request body
    var requestBody struct {
        Nama         string `json:"nama"`
        Harga        int    `json:"harga"`
        Stok         int    `json:"stok"`
        PerusahaanID string `json:"perusahaan_id"`
        KodeBarang   string `json:"kode"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // update existing barang in database
    existingBarang.Nama = requestBody.Nama
    existingBarang.Harga = requestBody.Harga
    existingBarang.Stok = requestBody.Stok
    existingBarang.PerusahaanID = requestBody.PerusahaanID
    existingBarang.KodeBarang = requestBody.KodeBarang
    result = initializers.DB.Save(&existingBarang)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to update barang",
            "data":    nil,
        })
        return
    }

    // return updated barang data
    data := gin.H{
        "id":            existingBarang.ID,
        "nama":          existingBarang.Nama,
        "harga":         existingBarang.Harga,
        "stok":          existingBarang.Stok,
        "kode":          existingBarang.KodeBarang,
        "perusahaan_id": existingBarang.PerusahaanID,
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Barang updated successfully",
        "data":    data,
    })
}

func UpdatePerusahaan(c *gin.Context) {
    // retrieve perusahaan ID from request
    perusahaanID := c.Param("id")

    // retrieve existing perusahaan from database
    var existingPerusahaan model.Perusahaan
    result := initializers.DB.First(&existingPerusahaan, perusahaanID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  "error",
            "message": "Perusahaan not found",
            "data":    nil,
        })
        return
    }

    // parse request body
    var requestBody struct {
        Nama     string `json:"nama"`
        Alamat   string `json:"alamat"`
        NoTelepon string `json:"no_telp"`
        KodePajak string `json:"kode"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // update existing perusahaan in database
    existingPerusahaan.Nama = requestBody.Nama
    existingPerusahaan.Alamat = requestBody.Alamat
    existingPerusahaan.NoTelepon = requestBody.NoTelepon
    existingPerusahaan.KodePajak = requestBody.KodePajak
    result = initializers.DB.Save(&existingPerusahaan)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to update perusahaan",
            "data":    nil,
        })
        return
    }

    // return updated perusahaan data
    data := gin.H{
        "id":      existingPerusahaan.ID,
        "nama":    existingPerusahaan.Nama,
        "alamat":  existingPerusahaan.Alamat,
        "no_telp": existingPerusahaan.NoTelepon,
        "kode":    existingPerusahaan.KodePajak,
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Perusahaan updated successfully",
        "data":    data,
    })
}