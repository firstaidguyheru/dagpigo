package dagpi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Client Class/Struct
type Client struct {
	auth *string
}

var c = Client{}

// Normal get request to get data with
func httpGet(url string) map[string]interface{} {
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

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)
	return data
}

// Attempting to get an image's buffer using a get request, can also be used as a normal get request
func getBuffer(url string) []byte {
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
// Data

func WTP() interface{} {
	// Returns an interface with all of the pokemon data
	data := httpGet("https://api.dagpi.xyz/data/wtp")
	return data
}

func Roast() interface{} {
	// Returns an interface containing a roast
	data := httpGet("https://api.dagpi.xyz/data/roast")
	roast := data["roast"]
	return roast
}

func Joke() interface{} {
	// Returns an interface containing a joke & id
	data := httpGet("https://api.dagpi.xyz/data/joke")
	return data
}

func Fact() interface{} {
	// Returns an interface containing a fact
	data := httpGet("https://api.dagpi.xyz/data/fact")
	fact := data["fact"]
	return fact
}

func Eightball() interface{} {
	// Returns an interface containing a response to 8ball question
	data := httpGet("https://api.dagpi.xyz/data/8ball")
	response := data["response"]
	return response
}

func Yomama() interface{} {
	// Returns an interface containing a description of yomama
	data := httpGet("https://api.dagpi.xyz/data/yomama")
	description := data["description"]
	return description
}

func RandomWaifu() interface{} {
	// Returns an interface containing data of a random waifu
	data := httpGet("https://api.dagpi.xyz/data/waifu")
	return data
}

func PickupLine() interface{} {
	// Returns an interface containing category & joke
	data := httpGet("https://api.dagpi.xyz/data/pickupline")
	return data
}

func Waifu(waifu_name string) interface{} {
	// Returns an interface containing data of a given waifu
	data := httpGet("https://api.dagpi.xyz/data/"+waifu_name)
	return data
}

func HeadLine() interface{} {
	// Returns an interface containing text and a bool, 'fake'
	data := httpGet("https://api.dagpi.xyz/data/headline")
	return data
}

func GTL() interface{} {
	// Returns an interface containing data of a random logo
	data := httpGet("https://api.dagpi.xyz/data/logo")
	return data
}

func Flag() interface{} {
	// Returns an interface containing data of a random flag
	data := httpGet("https://api.dagpi.xyz/data/flag")
	return data
}
// END

// Image Manipulation

func Pixelate(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/pixel/?url="+url)
	return buffer
}

func Colors(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/colors/?url="+url)
	return buffer
}

func America(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/america/?url="+url)
	return buffer
}

func Communism(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/communism/?url="+url)
	return buffer
}

func Triggered(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/triggered/?url="+url)
	return buffer
}

func Wasted(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/wasted/?url="+url)
	return buffer
}

func Invert(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/invert/?url="+url)
	return buffer
}

func Sobel(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/sobel/?url="+url)
	return buffer
}

func Hog(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/hog/?url="+url)
	return buffer
}

func Triangle(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/triangle/?url="+url)
	return buffer
}

func Blur(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/blur/?url="+url)
	return buffer
}

func RGB(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/rgb/?url="+url)
	return buffer
}

func Angel(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/angel/?url="+url)
	return buffer
}

func Satan(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/satan/?url="+url)
	return buffer
}

func Delete(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/delete/?url="+url)
	return buffer
}

func Fedora(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/fedora/?url="+url)
	return buffer
}

func Hitler(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/hitler/?url="+url)
	return buffer
}

func Wanted(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/wanted/?url="+url)
	return buffer
}

func Stringify(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/stringify/?url="+url)
	return buffer
}

func Mosiac(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/mosiac/?url="+url)
	return buffer
}

func Sithlord(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/sith/?url="+url)
	return buffer
}

func Jail(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/jail/?url="+url)
	return buffer
}

func Pride(url string, flag string) []byte {
	acceptableFlags := []string{
		"asexual", 
		"bisexual", 
		"gay", 
		"genderfluid",
		"genderqueer",
		"intersex",
		"lesbian",
		"nonbinary",
		"progress",
		"pan",
		"trans",
	}
	for _, acceptableFlag := range acceptableFlags {
		if acceptableFlag == flag {
			buffer := getBuffer("https://api.dagpi.xyz/image/pride/?url="+url+"&flag="+flag)
			return buffer
		} else {
			fmt.Println("That pride flag is not accepted")
		}
	}

	return nil
}

func Gay(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/gay/?url="+url)
	return buffer
}

func Trash(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/trash/?url="+url)
	return buffer
}

func Deepfry(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/deepfry/?url="+url)
	return buffer
}

func Ascii(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/ascii/?url="+url)
	return buffer
}

func Charcoal(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/ascii/?url="+url)
	return buffer
}

func Posterize(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/poster/?url="+url)
	return buffer
}

func Sepia(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/sepia/?url="+url)
	return buffer
}

func Swirl(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/swirl/?url="+url)
	return buffer
}

func Paint(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/paint/?url="+url)
	return buffer
}

func Night(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/night/?url="+url)
	return buffer
}

func Rainbow(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/rainbow/?url="+url)
	return buffer
}

func Magik(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/magik/?url="+url)
	return buffer
}

func FivegOneg(url1 string, url2 string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/5g1g/?url="+url1+"&url2="+url2)
	return buffer
}

func WhyAreYouGay(url1 string, url2 string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/whyareyougay/?url="+url1+"&url2="+url2)
	return buffer
}

func Obama(url1 string, url2 string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/obama/?url="+url1+"&url2="+url2)
	return buffer
}

func Tweet(url string, username string, text string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/tweet/?url="+url+"&username="+username+"&text="+text)
	return buffer
}

func Discord(url string, username string, text string, darkMode bool) []byte {
	if darkMode == true {
		buffer := getBuffer("https://api.dagpi.xyz/image/discord/?url="+url+"&username="+username+"&text="+text+"&dark="+"true")
		return buffer
	} else {
		buffer := getBuffer("https://api.dagpi.xyz/image/discord/?url="+url+"&username="+username+"&text="+text+"&dark="+"false")
		return buffer
	}
}

func YouTube(url string, username string, text string, darkMode bool) []byte {
	if darkMode == true {
		buffer := getBuffer("https://api.dagpi.xyz/image/yt/?url="+url+"&username="+username+"&text="+text+"&dark="+"true")
		return buffer
	} else {
		buffer := getBuffer("https://api.dagpi.xyz/image/yt/?url="+url+"&username="+username+"&text="+text+"&dark="+"false")
		return buffer
	}
}

func Retromeme(url string, topText string, bottomText string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/retromeme/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText)
	return buffer
}

func Motivational(url string, topText string, bottomText string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/motiv/?url="+url+"&top_text="+topText+"&bottom_text="+bottomText)
	return buffer
}

func Captcha(url string, text string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/captcha/?url"+url+"&text="+text)
	return buffer
}

func Modernmeme(url string, text string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/modernmeme/?url="+url+"&text="+text)
	return buffer
}

func Elmo(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/elmo/?url="+url)
	return buffer
}

func Rain(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/rain/?url="+url)
	return buffer
}

func Tv(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/tv/?url="+url)
	return buffer
}

func Album(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/allbum/?url="+url)
	return buffer
}

func StaticGlitch(url string) []byte {
	buffer := getBuffer("https://api.dagpi.xyz/image/static_glitch/?url="+url)
	return buffer
}
