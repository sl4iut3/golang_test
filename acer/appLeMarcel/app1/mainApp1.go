package main

import (
	"fmt"
	"sl/lemarcel"
)


func main() {
	lStations := lemarcel.Get_data_stations()
	for i := 0; i < len(lStations); i++ {
		fmt.Println(lStations[i].GetName(), lStations[i].GetListeVelo())
	}
}
