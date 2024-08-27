package lemarcel

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Tvelo struct {
	NumVelo string
	Station string
}

type Tstation struct {
	Name      string
	ListeVelo []string
}

func (s Tstation) GetName() string {
	return s.Name
}
func (s Tstation) GetListeVelo() []string {
	return s.ListeVelo
}

func traite_velos(stations []any, c chan []Tvelo) {
	fmt.Println("traite_velo : traitement de ", len(stations), " stations.")
	var lv []Tvelo
	for _, s := range stations {
		station := s.(map[string]any)
		name := station["name"].(string)
		docksData := (station["docks"].(map[string]any))["data"].([]any)
		for j := 0; j < len(docksData); j++ {
			d := docksData[j].(map[string]any)
			vaeNb := d["vehicule_number"]
			if vaeNb != nil {
				var veloR Tvelo
				veloR.Station = name
				x := vaeNb.(float64)
				veloR.NumVelo = strconv.FormatFloat(x, 'f', 0, 64)
				lv = append(lv, veloR)
			}
		}
		fmt.Println()
	}
	c <- lv
}

func traite_stations(stations []any, c chan []Tstation) {
	fmt.Println("traite_station : traitement de ", len(stations), " stations.")
	var ls []Tstation
	for _, s := range stations {
		var stationR Tstation
		station := s.(map[string]any)
		stationR.Name = station["name"].(string)
		docksData := (station["docks"].(map[string]any))["data"].([]any)
		for j := 0; j < len(docksData); j++ {
			d := docksData[j].(map[string]any)
			vaeNb := d["vehicule_number"]
			if vaeNb != nil {
				x := vaeNb.(float64)
				stationR.ListeVelo = append(stationR.ListeVelo, strconv.FormatFloat(x, 'f', 0, 64))
			}
		}
		fmt.Println()
		ls = append(ls, stationR)
	}
	c <- ls
}

func traitementStations(stations []any) ([]Tstation, []Tvelo) {
	cs := make(chan []Tstation)
	cv := make(chan []Tvelo)
	go traite_stations(stations, cs)
	go traite_velos(stations, cv)
	ls := <-cs
	lv := <-cv
	return ls, lv
}

func Get_data_stations() ([]Tstation, []Tvelo) {
	response, err := http.Get("https://api.cyclist.ecovelo.mobi/2021_12_07/stations?program=lemarcel&limit=30")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// remplacement de ioutil.ReadAll par io.ReadAll (maj go 1.16)
	//responseData, err := ioutil.ReadAll(response.Body)
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject map[string]any
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return traitementStations(responseObject["data"].([]any))
}

/* func main() {
	lStations := Get_data_stations()
	for i := 0; i < len(lStations); i++ {
		fmt.Println(lStations[i].Name, lStations[i].ListeVelo)
	} 
}*/
