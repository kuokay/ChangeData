package Model


import (
	"github.com/jinzhu/gorm"
	"time"
)



type User struct {
	gorm.Model
	Id      	int64  `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserName    string `gorm:"varchar(20);not null"`
	Email 		string `gorm:"varchar(50);not null;unique"`
	Password    string `gorm:"size:255;not null"`
	CreateTime  time.Time `gorm:"column:createtime;default:null" json:"createtime"`
}