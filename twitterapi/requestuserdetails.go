package twitterapi

import (
	"encoding/json"
	"fmt"

	//"image/jpeg"
	"net/http"
	"net/url"
	"os"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

type credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

//UserDetails contains selected USER Basic Information
type UserDetails struct {
	UserID       uint64 `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	Desc         string `json:"description"`
	Location     string `json:"location"`
	FriendsCount uint16 `json:"friends_count"`
	URL          string `json:"url"`
}

//UserInfo dfgsgs
var UserInfo UserDetails

func setCredentials(userToken *credentials, ID int) {
	userToken.ConsumerKey = "VTOzL1OQsK9QE0UWoh3SaVOmn"
	userToken.ConsumerSecret = "7j2ggtO3EVJBvMBlXSkDspm0cZF4d9NTthm5hOzsRXcdORLzwc"
	userToken.AccessToken = "1347416616-Jkvxh8jhNlFkAp54gzJTKAb2l6S8J2TUf4W3i4B"
	userToken.AccessTokenSecret = "lEKF0zRba81q98NzPC1q5u4TnUKEhb9S7rSCQ5Q1oZplO"
}

//returnClient() loads the credentials from struct to the api and returns a twittergo.Client object
func returnClient(userToken credentials) (client *twittergo.Client) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    userToken.ConsumerKey,
		ConsumerSecret: userToken.ConsumerSecret,
	}
	user := oauth1a.NewAuthorizedConfig(userToken.AccessToken, userToken.AccessTokenSecret)
	client = twittergo.NewClient(config, user)
	return
}

//returnClient() send a verification requesttwittergo.Client object
func verify(client *twittergo.Client) UserDetails {
	var (
		err  error
		req  *http.Request
		resp *twittergo.APIResponse
		//user     *twittergo.User
		endpoint string
		usrparam string
	)
	params := url.Values{}
	params.Add("include_email", "true")
	endpoint = "/1.1/account/verify_credentials.json?"
	usrparam = params.Encode()
	req, err = http.NewRequest("GET", (endpoint + usrparam), nil)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}

	resp, err = client.SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	tempResponse := resp.ReadBody()
	//fmt.Printf("\n\n%v\n\n", tempResponse)
	var jsonvar UserDetails
	json.Unmarshal([]byte(tempResponse), &jsonvar)

	return jsonvar

}

//RequestUserDetails loads the user tokes of USERID from a DB to a struct, uses the struct to authenticate client and uses client to access essential user details and store then to DB
func RequestUserDetails(ID int) {
	var (
		cred   credentials
		client *twittergo.Client
	)
	setCredentials(&cred, ID)
	//fmt.Printf("Credentials : \n%v\n\n", cred)
	client = returnClient(cred)

	UserInfo = verify(client)
	//fmt.Printf("User Details : \n%v\n\n", UserInfo)

}
