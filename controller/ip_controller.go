package controller

import (
	"encoding/json"
	"fmt"
	"ipProject/config"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func HandlePing(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong")

}
func HandleIpCount(w http.ResponseWriter, r *http.Request) {

	mockData := config.InputJson
	thresholdSet := os.Getenv("threshold")

	threshold, err := strconv.Atoi(thresholdSet)
	responseDomain := make([]string, 0)

	if err != nil {
		log.Printf("Error during converting string to int %v", err)
	}

	countOfActiveIpAddress := make(map[string]int)

	for _, data := range mockData.Data {
		countOfActiveIpAddress[data.HostName] = 0

	}

	for _, data := range mockData.Data {

		if data.Active {
			countOfActiveIpAddress[data.HostName]++
		}
	}

	for key, val := range countOfActiveIpAddress {

		if val <= threshold {
			responseDomain = append(responseDomain, key)

		}
	}

	sort.Strings(responseDomain)

	jsonBody, err := json.Marshal(responseDomain)
	if err != nil {
		log.Printf("Error during marshalling the array")
	}

	fmt.Fprintf(w, string(jsonBody))
	fmt.Println("Endpoint Hit: homePage", responseDomain)
}
