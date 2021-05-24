package dagpi

import (
	"net/http"
	"log"
	"io/ioutil"
)

// Client Class/Struct
type Client struct {
	auth *string
}

// Client constructor
func initClient(auth string) string {
	c := &Client{auth: &auth}
	return *c.auth
}

// Attempting to get an image's buffer using a get request
func getBuffer(url string, c Client) []byte {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Add("Authorization", *c.auth)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

