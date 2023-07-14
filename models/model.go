package model

import (
	// "gorm.io/gorm"
)

type User struct {
	ID       		uint   		`gorm:"primaryKey" json:"id"`
	Username 		string 		`gorm:"unique;not null" json:"username"`
	Password 		string 		`gorm:"not null" json:"-"`
	Name     		string 		`gorm:"not null" json:"name"`
	IsAdmin  		bool   		`gorm:"default:false" json:"is_admin"`
}

type Perusahaan struct {
	ID        		uint   		`gorm:"primaryKey" json:"id"`
	Nama      		string 		`gorm:"not null" json:"nama"`
	Alamat    		string 		`gorm:"not null" json:"alamat"`
	NoTelepon 		string 		`gorm:"not null" json:"no_telp"`
	KodePajak 		string 		`gorm:"unique;not null" json:"kode"`
	Barang   		[]Barang 	`gorm:"foreignKey:PerusahaanID" json:"barang"`
}

type Barang struct {
	ID           	uint   		`gorm:"primaryKey" json:"id"`
	Nama         	string 		`gorm:"not null" json:"nama"`
	Harga        	int    		`gorm:"not null;check:harga > 0" json:"harga"`
	Stok         	int    		`gorm:"not null;check:stok >= 0" json:"stok"`
	KodeBarang   	string 		`gorm:"unique;not null" json:"kode"`
	PerusahaanID 	string   	`gorm:"not null" json:"perusahaan_id"`
}

