package controllers

import (
	"singleservice/initializers"
	model "singleservice/models"
	"singleservice/auth"
	"net/http"
	"github.com/gin-gonic/gin"
	// "fmt"
)

func Login(c *gin.Context) {
    // get username and password from request body
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := c.BindJSON(&body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // authenticate user
	Users := []model.User{}
	initializers.DB.Where("username = ? AND password = ?", body.Username, body.Password).Find(&Users)
	// check if username and password match
	if len(Users) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Username and password do not match",
			"data":    nil,
		})
		return
	}

    // generate JWT token
    token, err := auth.GenerateToken(Users[0].ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to generate token",
            "data":    nil,
        })
        return
    }
    // return user data and token
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "User authenticated successfully",
        "data": gin.H{
            "user": gin.H{
                "username": Users[0].Username,
                "name":     Users[0].Name,
            },
            "token": token,
        },
    })
	return;
}

func Register(c *gin.Context) {
    // get username and password from request body
    var body struct {
        Username    string `json:"username"`
        Password    string `json:"password"`
        Name        string `json:"name"`
    }

    err := c.BindJSON(&body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // create new user in database
    newUser := model.User{
        Username: body.Username,
        Password: body.Password,
        Name:     body.Name,
    }
    result := initializers.DB.Create(&newUser)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to create user",
            "data":    nil,
        })
        return
    }

    // return created user data
    data := gin.H{
        "id":       newUser.ID,
        "username": newUser.Username,
        "name":     newUser.Name,
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "User created successfully",
        "data":    data,
    })
}

func CreateBarang(c *gin.Context) {
    // parse request body
    var requestBody struct {
        Nama         string `json:"nama" binding:"required"`
        Harga        int    `json:"harga" binding:"required"`
        Stok         int    `json:"stok" binding:"required"`
        PerusahaanID string  `json:"perusahaan_id" binding:"required"`
        KodeBarang   string `json:"kode" binding:"required"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // create new barang in database
    barang := model.Barang{
        Nama:         requestBody.Nama,
        Harga:        requestBody.Harga,
        Stok:         requestBody.Stok,
        PerusahaanID: requestBody.PerusahaanID,
        KodeBarang:   requestBody.KodeBarang,
    }
    result := initializers.DB.Create(&barang)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to create barang",
            "data":    nil,
        })
        return
    }

    // return created barang data
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
        "message": "Barang created successfully",
        "data":    data,
    })
}

func CreatePerusahaan(c *gin.Context) {
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

    // create new perusahaan in database
    newPerusahaan := model.Perusahaan{
        Nama:     requestBody.Nama,
        Alamat:   requestBody.Alamat,
        NoTelepon: requestBody.NoTelepon,
        KodePajak: requestBody.KodePajak,
    }
    result := initializers.DB.Create(&newPerusahaan)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to create perusahaan",
            "data":    nil,
        })
        return
    }

    // return created perusahaan data
    data := gin.H{
        "id":      newPerusahaan.ID,
        "nama":    newPerusahaan.Nama,
        "alamat":  newPerusahaan.Alamat,
        "no_telp": newPerusahaan.NoTelepon,
        "kode":    newPerusahaan.KodePajak,
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Perusahaan created successfully",
        "data":    data,
    })
}

func CreateTransaksi(c *gin.Context){
    var requestBody struct {
        IDPembeli     uint   `json:"id_pembeli"`
        NamaBarang    string `json:"nama_barang"`
        TotalHarga    int    `json:"total_harga"`
    }

    err := c.BindJSON(&requestBody)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "data":    nil,
        })
        return
    }

    // create new transaksi in database
    newTransaksi := model.Transaksi{
        IDPembeli:     requestBody.IDPembeli,
        NamaBarang:    requestBody.NamaBarang,
        TotalHarga:    requestBody.TotalHarga,
    }

    result := initializers.DB.Create(&newTransaksi)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to create transaksi",
            "data":    nil,
        })
        return
    }

    // return created transaksi data
    data := gin.H{
        "id":           newTransaksi.ID,
        "id_pembeli":   newTransaksi.IDPembeli,
        "nama_barang":  newTransaksi.NamaBarang,
        "total_harga":  newTransaksi.TotalHarga,
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Transaksi created successfully",
        "data":    data,
    })
}