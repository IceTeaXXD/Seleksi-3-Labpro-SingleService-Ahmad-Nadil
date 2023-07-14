package controllers

import (
	"singleservice/initializers"
	model "singleservice/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func DeleteBarang(c *gin.Context) {
    // retrieve barang ID from request
    barangID := c.Param("id")

    // delete existing barang from database
    result := initializers.DB.Delete(&model.Barang{}, barangID)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to delete barang",
            "data":    nil,
        })
        return
    }

    // return deleted barang data
    data := gin.H{
        "id":            barangID,
        "nama":          "",
        "harga":         0,
        "stok":          0,
        "kode":          "",
        "perusahaan_id": "",
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Barang deleted successfully",
        "data":    data,
    })
}

func DeletePerusahaan(c *gin.Context) {
    // retrieve perusahaan ID from request
    perusahaanID := c.Param("id")

    // delete perusahaan from database
    result := initializers.DB.Delete(&model.Perusahaan{}, perusahaanID)
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  "error",
            "message": "Perusahaan not found",
            "data":    nil,
        })
        return
    }
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to delete perusahaan",
            "data":    nil,
        })
        return
    }

    // return deleted perusahaan data
    data := gin.H{
        "id":      perusahaanID,
        "nama":    "",
        "alamat":  "",
        "no_telp": "",
        "kode":    "",
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Perusahaan deleted successfully",
        "data":    data,
    })
}