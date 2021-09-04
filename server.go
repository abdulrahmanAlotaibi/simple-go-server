package main


import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)


func main (){
	
	var num = rand.Intn(10) + 1
    fmt.Println(num)

	fmt.Println(18.91 * 0.543218192)
	http.HandleFunc("/_health", func(res http.ResponseWriter, req *http.Request){
		fmt.Println(req)
		fmt.Fprintf(res,"hello!")


		if req.Method == "GET" {
			fmt.Fprintf(res,"hello!")
		}

		if req.Method == "POST" {
			fmt.Fprintf(res,"THIS IS POST!")
		}
	})
	
	
	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}