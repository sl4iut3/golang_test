package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
  "sl/lemarcel"
)

func main() {
	app := tview.NewApplication()

  title := tview.NewTextView().
  SetTextAlign(tview.AlignCenter).SetDynamicColors(true).SetText("[red]Marcel[white] v0.9")
  footer := tview.NewTextView().
  SetTextAlign(tview.AlignLeft).SetDynamicColors(true).SetText("[red]copyright 2024[white] v0.9")

	table := tview.NewTable().
  SetBorders(true)
	lStations := lemarcel.Get_data_stations()
	//	lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
	cols, rows := 4, len(lStations)/2+1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ind := r + c/2*rows
			color := tcell.ColorWhite
			if c%2 == 0 {
				color = tcell.ColorYellow
			}
			text := " "
			if ind < len(lStations) {
				switch c {
				case 0:
					text = lStations[ind].GetName()
				case 2:
					text = lStations[ind].GetName()
				case 1:
					text = strings.Join(lStations[ind].GetListeVelo(), " ")
				case 3:
					text = strings.Join(lStations[ind].GetListeVelo(), " ")
				}
			}
			table.SetCell(r, c,
				tview.NewTableCell(text).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
		}
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape || key == 'q' {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

  grid := tview.NewGrid().SetRows(3,0,3).SetColumns(10,-8,10).SetBorders(true).
              AddItem(title,0,0,1,3,0,0,false).
                AddItem(table,1,1,1,1,0,0,true).
                AddItem(footer,2,0,1,3,0,0,false)

	//if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
