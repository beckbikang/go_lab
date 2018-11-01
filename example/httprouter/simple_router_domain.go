package main


import (

	"github.com/beckbikang/httprouter"
	"net/http"
	"log"
	"fmt"
)

func Index2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("do hello handle")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type HostSwitch map[string]http.Handler

// Implement the ServeHTTP method on our new type
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		// Handle host names for which no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func main() {

	//支持多个网站
	//初始化路由器
	// Initialize a router as usual
	router := httprouter.New()
	router.GET("/", Index2)
	router.GET("/hello/:name", Hello2)

	// Make a new HostSwitch and insert the router (our http handler)
	// for example.com and port 12345
	//定义处理器
	hs := make(HostSwitch)
	hs["127.0.0.1:8881"] = router

	// Use the HostSwitch to listen and serve on port 12345
	log.Fatal(http.ListenAndServe(":12345", hs))
}
