package twitterapi

import (
	"fmt"

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
	fmt.Println(cred)
	return cred
}

func readUserDBtostruct(db *gorm.DB, name string) User {
	var user User
	db.Raw("SELECT * FROM User  WHERE Accname = ?", name).Scan(&user)
	//fmt.Println(user)
	return user
}

//Operation exported function
func Operation() {
	db, err := gorm.Open("sqlite3", "twitter1.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)
	userinfo := readUserDBtostruct(db, "Parikshit Ghosh")
	usercred := readCredDBtostruct(db, userinfo.ID)
	//userinfo, usercred := readDBtostruct(db)

	fmt.Println("\n\nBaidurya details :")
	fmt.Printf("\tApp ID : %d\n", userinfo.ID)
	fmt.Printf("\tApp Name %d\n", userinfo.Appid)
	fmt.Printf("\tApp Name %s\n", userinfo.Appname)
	fmt.Printf("\tApp Name %s\n", userinfo.Accname)
	fmt.Println("\n\nBaidurya Credentials :")
	fmt.Printf("\tAPI Key %s\n", usercred.Apikey)
	fmt.Printf("\tAPI Secret %s\n", usercred.Apisecret)
	fmt.Printf("\tAccess Key %s\n", usercred.Accesskey)
	fmt.Printf("\tAccess Secret %s\n", usercred.Accesssecret)

}
