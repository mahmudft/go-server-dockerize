package server

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)


func main() {
	port = flag.int("port", 2222, "Specify a port for your web server")
	workdir = flag.string("dir", ".", "specify where to run server")
	flag.Parse()

	http.Handle('/', http.FileServer(http.Dir(*workdir)))
	fmt.Printf("Server running on port : %d ", *port)
	log.Fatal(http.ListenAndServe(":"+*p, nil))

}