package opts

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", Cfg.App.Db.UserName, Cfg.App.Db.Password, "tcp", Cfg.App.Db.Url, Cfg.App.Db.Port, Cfg.App.Db.StoreHouse)

	DB, err = gorm.Open(Cfg.App.Db.Type, conn)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	DB.SingularTable(true)
}

type User struct {
	ID           int    `json:"id",gorm:"primary_key"`
	Email        string `json:"email" binding:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Roleid       string `json:"roleid"`
}

func (user *User) QueryByEntry(userName string, password string) (User, error) {
	e := DB.Find(user, "username=? and password=?", userName, password).Error
	return *user, e
}

func (user *User) QueryByToken(token string) (User, error) {
	e := DB.Find(user, "token=?", token).Error
	return *user, e
}

func (user *User) UpdateTokenByUser() {
	update := DB.Model(&user).Where("id = ?", user.ID).Update("token", user.Token)
	println(update)
}

func (user *User) ClearTokenByUser() {
	DB.Model(&user).Update("token", "")
}
