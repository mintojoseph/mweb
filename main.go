package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

//Env lookup
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//Camel to Space
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func camelToSpaced(str string) string {
	spaced := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	spaced = matchAllCap.ReplaceAllString(spaced, "${1} ${2}")
	return (spaced)
}

//helloworld HTTP Handler
func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("name")
	if len(value) > 0 {
		fmt.Fprintf(w, "Hello %s\n", camelToSpaced(value))
	} else {

		fmt.Fprintf(w, "Hello Stranger\n")
	}
}

func main() {

	port := "8080"
	portNo := getEnv("PORT", port)
	fmt.Println("PORT:", portNo)

	http.HandleFunc("/helloworld", helloworldHandler)

	log.Fatal(http.ListenAndServe(":"+portNo, nil))
}
