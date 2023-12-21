package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("fetching http://192.168.0.1/\n")
	fetch(w, r, "http://192.168.0.1/")

	fmt.Fprintf(w, "\n")

	// fmt.Printf("fetching http://127.0.0.1/")
	// fetch(w, r, "http://127.0.0.1/") //todo catch errors
}
func fetch(w http.ResponseWriter, r *http.Request, urlString string) {
	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "%+v\n", *resp.Request)

	URL := resp.Request.URL.String()
	m := getMapParameters(URL)
	finalURL := generateNewUrl(m)

	_, err = http.Get(finalURL)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "The URL you ended up at is: %v\n", finalURL)
}
func generateNewUrl(reeplacebles map[string]string) string {
	finalURL := "https://portal-wifi-wag.telecentro.net.ar/VGVsZUNlbnRybyBXaWZp/init_process?client_mac=" +
		reeplacebles["client_mac"] +
		"&mac=" + reeplacebles["client_mac"] +
		"&ssid=TeleCentro%20Wifi&nas_ip=" + reeplacebles["nas_ip"] +
		"&hash=" + reeplacebles["hash"] + "&tipo=outdoor"
	return finalURL
}
func getMapParameters(finalURL string) map[string]string {
	parameters := [4]string{"client_mac", "mac", "nas_ip", "hash"}
	split := strings.Split(finalURL, "&")
	m := make(map[string]string)

	for _, element := range split {
		mapSplit := strings.Split(element, "=")
		for _, parameter := range parameters {
			if strings.Contains(element, parameter) {
				m[parameter] = mapSplit[1]
			}
		}
	}
	return m
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
