package main

import (
	"io/ioutil"
	AppConfig "main/appconfig"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	// setup http router
	router := httprouter.New()
	router.GET(AppConfig.Get("local_uri"), getRequest)
	log.Fatal(http.ListenAndServe(AppConfig.Get("listen"), router))
}

// get basic http request
func getRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	key := queryValues.Get(AppConfig.Get("parameter_name"))
	if len(key) == 0 {
		log.Warn("Got request with empty key")
		http.NotFound(w, r)
		return
	}
	log.Info("Token requested: " + key)
	sendRequest(key, w, r)

}

// send http request to other web server
func sendRequest(key string, w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: 1 * time.Second}
	res, err := client.Get(AppConfig.Get("remote_uri") + key)
	if err != nil {
		http.NotFound(w, r)
		log.Error("Got error upon request: " + err.Error())
		return
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.NotFound(w, r)
		log.Error("Got error while streaming request: " + err.Error())
		return
	}
	if res.StatusCode != 200 {
		http.NotFound(w, r)
		log.Error("Got error from gateway:" + string(bytes))
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		http.NotFound(w, r)
		log.Error("Got error while send response: " + err.Error())
		return
	}
}
