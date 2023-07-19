package route

import (
	"ipProject/controller"
	"log"
	"net/http"
	"os"
)

func InitRoutes() {

	http.HandleFunc("/",controller.HandlePing)
	http.HandleFunc("/count/domain",controller.HandleIpCount)


	if os.Getenv("local") != "1"{
		log.Fatal(http.ListenAndServe(":10000", nil))

	}
}