package twitterapi

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //using sqlite
)

//Userprofile struct stores user credentials
type Userprofile struct {
	ID            int64  `gorm:"column:ID"`
	Username      string `gorm:"column:Username"`
	Fullname      string `gorm:"column:Fullname"`
	Emailid       string `gorm:"column:Emailid"`
	Gender        string `gorm:"column:Gender"`
	URL           string `gorm:"column:Url"`
	Status        string `gorm:"column:Status"`
	Location      string `gorm:"column:Location"`
	Twitterlinked string `gorm:"column:Twitterlinked"`
}

//Twittercred store user tokens
type Twittercred struct {
	Userid       int    `gorm:"column:Userid"`
	Username     string `gorm:"column:Username"`
	Apikey       string `gorm:"column:Apikey"`
	Apisecret    string `gorm:"column:Apisecret"`
	Accesskey    string `gorm:"column:Accesskey"`
	Accesssecret string `gorm:"column:Accesssecret"`
	Appid        uint32 `gorm:"column:Appid"`
	Appname      string `gorm:"column:Appname"`
}

//UpdateUserProfile creates a new user through registration
func UpdateUserProfile(newDetails Userprofile) {
	db, err := gorm.Open("sqlite3", "site.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	if newDetails.Fullname != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("Fullname", newDetails.Fullname)
	}

	if newDetails.Emailid != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("Emailid", newDetails.Emailid)
	}
	if newDetails.Gender != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("Gender", newDetails.Gender)
	}
	if newDetails.URL != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("URL", newDetails.URL)
	}
	if newDetails.Status != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("Status", newDetails.Status)
	}
	if newDetails.Location != "" {
		db.Model(&Userprofile{}).Where("Username = ?", newDetails.Username).Update("Location", newDetails.Location)
	}
}

//CreateNewUserProfile creates a new user through registration
func CreateNewUserProfile(newUser Userprofile) {
	db, err := gorm.Open("sqlite3", "site.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	db.Select("Username", "Fullname", "Emailid", "Gender", "URL", "Status", "Location", "Twitterlinked").Create(&newUser)
}

//Viewprofile sends you your profile details with associated username otherwise returns error!
func Viewprofile(uname string) (Userprofile, int) {
	db, err := gorm.Open("sqlite3", "site.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	var (
		errno int
		info  Userprofile
	)

	result := db.Raw("SELECT * FROM Userprofiles  WHERE Username = ?", uname).Scan(&info)
	flag := result.RecordNotFound()
	if flag {
		errno = 1
	} else {
		errno = 0
	}
	return info, errno
}

//Removeprofile deletes User with specified User.Username
func Removeprofile(uname string) error {
	db, err := gorm.Open("sqlite3", "site.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	query := db.Where("Username = ?", uname).Delete(&Userprofile{})

	return query.Error
}

//TwitterCredbyUName exports Twittercred struct of UserName passed
func TwitterCredbyUName(uname string) (Twittercred, int) {
	db, err := gorm.Open("sqlite3", "site.db")
	if err != nil {
		panic("can't connect to database")
	}
	defer db.Close()
	db.LogMode(true)

	var (
		errno int
		cred  Twittercred
	)

	result := db.Raw("SELECT * FROM Twittercreds  WHERE Username = ?", uname).Scan(&cred)
	flag := result.RecordNotFound()
	if flag {
		errno = 1
	} else {
		errno = 0
	}
	return cred, errno
}
