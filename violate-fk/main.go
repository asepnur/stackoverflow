package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	ID       string  `gorm:"column:id"`
	ClientID string  `gorm:"type:varchar(255);primaryKey;column:client_id"`
	Name     string  `gorm:"column:name"`
	Albums   []Album `gorm:"foreignKey:request_client_id;references:client_id"`
}

type Album struct {
	ID              string    `gorm:"type:varchar(255);primaryKey;column:id"`
	RequestClientID string    `gorm:"type:varchar(255);column:request_client_id"`
	Pictures        []Picture `gorm:"foreignKey:album_id;references:id"`
}

type Picture struct {
	PictureID   string `gorm:"primaryKey;column:picture_id"`
	AlbumID     string `gorm:"type:varchar(255);column:album_id"`
	Description string `gorm:"column:description"`
}

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/violate?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := gormDB.AutoMigrate(&Order{}, &Album{}, &Picture{}); err != nil {
		log.Println("err: ", err)
		return
	}
	picture := []Picture{
		{
			PictureID:   "pic_1",
			Description: "test pic",
		},
	}
	albums := []Album{
		{
			ID:              "al_1",
			Pictures:        picture,
			RequestClientID: "",
		},
	}
	orders := Order{
		ID:       "abc",
		ClientID: "client1",
		Name:     "Roy",
	}
	if err := gormDB.Save(orders).Error; err != nil {
		return
	}
	if err := gormDB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&orders).Association("Albums").Append(albums); err != nil {
		return
	}

}
