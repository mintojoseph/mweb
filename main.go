package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Env lookup
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var port = "8080"
var portNo = getEnv("HTTPPORT", port)

var portHTTP = flag.String("port", portNo, "HTTP port. HTTPPORT environment variable can also be used.")

// Camel to Space
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func camelToSpaced(str string) string {
	spaced := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	spaced = matchAllCap.ReplaceAllString(spaced, "${1} ${2}")
	return (spaced)
}

// helloworld HTTP Handler
func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("name")
	if len(value) > 0 {
		fmt.Fprintf(w, "Hello %s\n", camelToSpaced(value))
	} else {

		fmt.Fprintf(w, "Hello Stranger\n")
	}
}

// GitCommit for hash
var GitCommit string

func versionzHandler(w http.ResponseWriter, r *http.Request) {

	type versionDetails struct {
		Name string
		Hash string
	}
	basket := versionDetails{
		Name: "mweb",
		Hash: GitCommit,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonData))
}

// Logger logs incoming requests, including response status.
func Logger(out io.Writer, h http.Handler) http.Handler {
	logger := log.New(out, "", 0)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &responseObserver{ResponseWriter: w}
		h.ServeHTTP(o, r)
		addr := r.RemoteAddr
		if i := strings.LastIndex(addr, ":"); i != -1 {
			addr = addr[:i]
		}
		logger.Printf("%s %d %q",

			time.Now().Format(time.RFC3339),
			o.status,
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
		)
	})
}

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func main() {
	flag.Parse()

	fmt.Println("HTTP PORT:", *portHTTP)

	http.HandleFunc("/helloworld", helloworldHandler)
	http.HandleFunc("/versionz", versionzHandler)
	log.Fatal(http.ListenAndServe(":"+*portHTTP, Logger(os.Stderr, http.DefaultServeMux)))
}
