package main

import (
	"fmt"
	"os"
	"sl/lemarcel"
	"text/template"
)

func main() {
	const form = `station {{.GetName}} : {{"\t\t\t\t\t"}} {{.ListeVelo}} {{"\n"}}`
	t := template.Must(template.New("form").Parse(form))

	// var lStations [] lemarcel.Tstation
	lStations, lVelos := lemarcel.Get_data_stations()

	for _, r := range lStations {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			os.Exit(1)
		}
	}
	for _, r := range lVelos {
		fmt.Println(r.NumVelo, r.Station)
	}
}
