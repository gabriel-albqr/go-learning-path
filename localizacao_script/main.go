package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Query      string  `json:"query"` // IP
}

func main() {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var loc Location
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		panic(err)
	}

	fmt.Println("IP:", loc.Query)
	fmt.Println("País:", loc.Country)
	fmt.Println("Estado:", loc.RegionName)
	fmt.Println("Cidade:", loc.City)
	fmt.Printf("Latitude: %.5f\n", loc.Lat)
	fmt.Printf("Longitude: %.5f\n", loc.Lon)
}
