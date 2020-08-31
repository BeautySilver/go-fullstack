package api

import (
	"../auto"
	"../config"
	"./router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("Running on %d...", config.PORT)
	listen(config.PORT)

}


func listen(port int){
	r:=router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.LoadCORS(r)))
}