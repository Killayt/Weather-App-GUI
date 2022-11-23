package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Killayt/Weather-App-GUI/iternal"
)

func main() {
	// making window
	a := app.New()
	w := a.NewWindow("Weather K")
	w.Resize(fyne.Size{Width: 500, Height: 700})

	// Creating input form
	entryToFind := widget.NewEntry()   // Creating input form
	entryToFind.SetPlaceHolder("City") // placeholder "City"

	// Getting data value from button btn
	answare := widget.NewLabel("")

	// Button Find and handler
	btn := widget.NewButton("Find", func() { // Handler func

		city := entryToFind.Text // getting inputed text

		//  city
		data, err := iternal.Target(city)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(data)

		// Parse the data and display it in the window
		formatToString := strconv.FormatFloat(data.Main.Celsius, 'g', 4, 64)
		result := "Today in " + data.Name + "\n" + formatToString + " celsius."
		answare.SetText(result)

	})

	w.SetContent(
		container.NewVBox(
			entryToFind,
			btn,
			answare,
		))

	w.ShowAndRun()
}
