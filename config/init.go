package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ipProject/model"
	"os"
)
var (

	InputJson model.IpConfig
)

// Init the services for config
func Init() {

	MockIpData()
	SetThreshold()

}

// Set the environment variable for threshod given
func SetThreshold(){

	os.Setenv("threshold", "1")
	os.Setenv("local", "0")
}

// Read ip related inout data from the json file
func MockIpData(){

	jsonFile, err := os.Open("etc/mock.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened mock.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &InputJson)

	fmt.Println(InputJson)
}
