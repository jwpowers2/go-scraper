package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Address struct {
	Address   string `json: "address"`
	Store     string `json: "store"`
	Thumb     string `json: "thumb"`
	Id        string `json: "id"`
	Distance  string `json: "distance"`
	Permalink string `json: "permalink"`
	Address2  string `json: "address2"`
	City      string `json: "city"`
	State     string `json: "state"`
	Zip       string `json: "zip"`
	Country   string `json: "country"`
	Lat       string `json: "lat"`
	Lng       string `json: "lng"`
	Phone     string `json: "phone"`
	Fax       string `json: "fax"`
	Email     string `json: "email"`
	//Hours             string `json: "hours"`
	Url               string `json: "url"`
	CategoryMarkerUrl string `json: "categoryMarkerUrl"`
	Terms             string `json: "terms"`
}

type LatLong struct {
	Lat  string `json: "lat"`
	Long string `json: "long"`
}

func main() {

	latLongArray := []LatLong{}
	latLongArray = append(latLongArray, LatLong{Lat: "35.89023", Long: "-78.91751"})
	latLongArray = append(latLongArray, LatLong{Lat: "35.99812", Long: "-78.89204"})
	latLongArray = append(latLongArray, LatLong{Lat: "35.0525449", Long: "-78.878322"})
	addressList := []Address{}

	c := colly.NewCollector(
		colly.AllowedDomains("cookout.com"),
	)

	c.OnResponse(func(r *colly.Response) {

		var tmpAddresses []Address
		err := json.Unmarshal(r.Body, &tmpAddresses)
		if err != nil {
			fmt.Println(err)
		}
		addressList = append(addressList, tmpAddresses...)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	for _, s := range latLongArray {
		url := fmt.Sprintf("https://cookout.com/wp-admin/admin-ajax.php?action=store_search&lat=%v&lng=%v&max_results=50&search_radius=10", s.Lat, s.Long)
		c.Visit(url)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(addressList)
}
