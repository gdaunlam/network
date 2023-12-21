package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	googleURL     = "https://www.google.com"
	retryInterval = 5 * time.Second
	urlString = "http://192.168.0.1/"
)

func fetch() error {
	resp, err := http.Get(urlString)
	if err != nil { return err }
	fmt.Printf("%+v\n", *resp.Request)

	URL := resp.Request.URL.String()
	m := getMapParameters(URL)
	finalURL := generateNewUrl(m)

	_, err = http.Get(finalURL)
	if err != nil { return err }
	fmt.Printf("The URL you ended up at is: %v\n", finalURL)
	return nil;
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
	for {
		err := fetch()
		if err != nil { fmt.Printf("Unlock Error: %v.\n", err) }
		time.Sleep(retryInterval)
	}
}