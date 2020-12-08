package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sellerapps/internal/models"
	"strings"
)

//collection := helper.ConnectDB()
func homePage(w http.ResponseWriter, r *http.Request) {

	//check if post method
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	//scrape the url page
	var product models.Product
	var scerr error
	var MURLs []string = r.URL.Query()["url"]
	if len(MURLs) <= 0 {
		errorsend := Error{
			ErrorInfo: "url parameter not present",
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorsend)
		return
	}
	MURL := MURLs[0]
	product, scerr = scrape(MURL)
	// if error in scraping
	if scerr != nil {
		errorsend := Error{
			ErrorInfo: scerr.Error(),
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorsend)
		return
	}

	var productDataReq models.ProductData = models.ProductData{
		URL:     MURL,
		Product: &product,
	}
	//jsonify data
	jsonReq, err := json.Marshal(productDataReq)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		errorsend := Error{
			ErrorInfo: err.Error(),
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errorsend)
		return
	}

	//send post request
	response, err := http.Post("http://localhost:8888/", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	// if error in post request
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		errorsend := Error{
			ErrorInfo: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errorsend)
		return
	} else {

		bodyBytes, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(bodyBytes))

		//for if not able to connect to mongodb or any other error
		if strings.Contains(string(bodyBytes), "error_info") {
			var errorsend Error
			json.Unmarshal(bodyBytes, &errorsend)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(errorsend)
			return
		}
		var sendProdInfo models.ProductData
		json.Unmarshal(bodyBytes, &sendProdInfo)
		json.NewEncoder(w).Encode(sendProdInfo)
	}

	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

type Error struct {
	ErrorInfo string `json:"error_info,omitempty" bson:"error_info,omitempty"`
}
