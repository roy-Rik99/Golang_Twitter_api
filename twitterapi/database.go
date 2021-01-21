package twitterapi

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //using sqlite
)

//User struct stores user credentials
type User struct {
	ID      int    `gorm:"column:ID"`
	Appid   uint32 `gorm:"column:Appid"`
	Appname string `gorm:"column:Appname"`
	Accname string `gorm:"column:Accname"`
}

//Cred store user tokens
type Cred struct {
	ID           int    `gorm:"column:ID"`
	Apikey       string `gorm:"column:Apikey"`
	Apisecret    string `gorm:"column:Apisecret"`
	Accesskey    string `gorm:"column:Accesskey"`
	Accesssecret string `gorm:"column:Accesssecret"`
}

func readCredDBtostruct(db *gorm.DB, id int) Cred {

	var cred Cred
	db.Raw("SELECT * FROM Cred  WHERE ID = ?", id).Scan(&cred)
	//fmt.Println(cred)
	return cred
}
func readUserDBtostruct(db *gorm.DB, name string) User {
	var user User
	db.Raw("SELECT * FROM User  WHERE Accname = ?", name).Scan(&user)
	//fmt.Println(user)
	return user
}

//ObtainTokenbyName exports Cred struct of UserName Passed
func ObtainTokenbyName(n string) Cred {
	db, err := gorm.Open("sqlite3", "twitter.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)
	userinfo := readUserDBtostruct(db, n)
	usercred := readCredDBtostruct(db, userinfo.ID)
	return usercred
}
