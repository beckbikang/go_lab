package main



import (
	"fmt"
	"github.com/beckbikang/httprouter"
	"net/http"
	"log"
)


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("do hello handle")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/hello2/:name", Hello)
	//why error
	//router.GET("/hello/tt/:cname2", Hello)
	log.Fatal(http.ListenAndServe(":8089", router))
}