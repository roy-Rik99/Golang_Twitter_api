package twitterapi

import (
	"encoding/json"
	"fmt"
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

//Twitterinfo contains selected USER Basic Information
type Twitterinfo struct {
	TwitterID  int64  `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Desc       string `json:"description"`
	Location   string `json:"location"`
	URL        string `json:"url"`
}

//returnClient() loads the credentials from struct to the api and returns a twittergo.Client object
func returnClient(userToken Twittercred) (client *twittergo.Client) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    userToken.Apikey,
		ConsumerSecret: userToken.Apisecret,
	}
	user := oauth1a.NewAuthorizedConfig(userToken.Accesskey, userToken.Accesssecret)
	client = twittergo.NewClient(config, user)
	return
}

//returnClient() send a verification requesttwittergo.Client object
func verify(client *twittergo.Client) Twitterinfo {
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
	var jsonvar Twitterinfo
	json.Unmarshal([]byte(tempResponse), &jsonvar)

	return jsonvar

}

//RequestUserDetails loads the user tokes of USERName from a DB to a struct, uses the struct to authenticate client and uses client to access essential user details and store then to DB
func RequestUserDetails(cred Twittercred) Twitterinfo {
	client := returnClient(cred)
	//UserInfo obtains Twitter User Information in JSON format
	var twitterinfo Twitterinfo
	twitterinfo = verify(client)
	return twitterinfo
}
