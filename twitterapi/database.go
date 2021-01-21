package twitterapi

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //using sqlite
)

//User struct stores user credentials
type User struct {
	ID      int
	Appid   uint32
	Appname string
	Accname string
}

//Cred store user tokens
type Cred struct {
	ID           int
	Apikey       string
	Apisecret    string
	Accesskey    string
	Accesssecret string
}

func readDBtostruct(db *gorm.DB) (User, Cred) {

	var user User
	var cred Cred

	//b.First(&user, "ID = ?", 69)
	//db.First(&cred, "ID = ?", 69)
	db.Raw("SELECT * FROM User  WHERE ID = ?", 69).Scan(&user)
	db.Raw("SELECT * FROM Cred  WHERE ID = ?", 69).Scan(&cred)
	fmt.Println(user)
	fmt.Println(cred)
	return user, cred
}

//Operation exported function
func Operation() {
	db, err := gorm.Open("sqlite3", "twitter.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)
	readDBtostruct(db)
	/*userinfo, usercred := readDBtostruct(db)

	fmt.Println("\n\nBaidurya details :")
	fmt.Printf("\tApp ID : %v\n", userinfo.ID)
	fmt.Printf("\tApp Name %v\n", userinfo.Appid)
	fmt.Printf("\tApp Name %v\n", userinfo.Appname)
	fmt.Printf("\tApp Name %v\n", userinfo.Accname)
	fmt.Println("\n\nBaidurya Credentials :")
	fmt.Printf("\tApp ID : %v\n", usercred.ID)
	fmt.Printf("\tAPI Key %v\n", usercred.Apikey)
	fmt.Printf("\tAPI Secret %v\n", usercred.Apisecret)
	fmt.Printf("\tAccess Key %v\n", usercred.Accesskey)
	fmt.Printf("\tAccess Secret %v\n", usercred.Accesssecret)*/

}
