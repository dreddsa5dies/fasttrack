package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sort"
)

type dataJSON struct {
	Offers []struct {
		OfferID   string `json:"offer_id"`
		MarketSku int    `json:"market_sku"`
		Price     int    `json:"price"`
	} `json:"offers"`
}

type expJSON struct {
	Offers []struct {
		MarketSku int    `json:"market_sku"`
		OfferID   string `json:"offer_id"`
		Price     int    `json:"price"`
	} `json:"offers"`
}

type Offers struct {
	MarketSku int    `json:"market_sku"`
	OfferID   string `json:"offer_id"`
	Price     int    `json:"price"`
}

func (a expJSON) Len() int { return len(a.Offers) }
func (a expJSON) Swap(i, j int) {
	a.Offers[i].Price, a.Offers[j].Price = a.Offers[j].Price, a.Offers[i].Price
}
func (a expJSON) Less(i, j int) bool { return a.Offers[i].Price < a.Offers[j].Price }

func main() {
	var data []dataJSON

	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var temp dataJSON
		err := json.Unmarshal(scanner.Bytes(), &temp)
		if err != nil {
			log.Println(err)
		}
		data = append(data, temp)
	}

	var alldata dataJSON
	for k := range data {
		t := data[k]
		alldata.Offers = append(alldata.Offers, t.Offers...)
	}

	var exp expJSON

	for k := range alldata.Offers {
		tmp := Offers{}
		tmp.MarketSku = alldata.Offers[k].MarketSku
		tmp.OfferID = alldata.Offers[k].OfferID
		tmp.Price = alldata.Offers[k].Price
		exp.Offers = append(exp.Offers, tmp)
	}

	sort.Sort(exp)

	expJSON, err := json.Marshal(exp)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(string(expJSON))
	if err != nil {
		log.Fatal(err)
	}
}
