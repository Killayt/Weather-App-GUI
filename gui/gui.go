package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Killayt/Weather-App-GUI/config"
)

func MakeGui() {
	// making window
	a := app.New()
	w := a.NewWindow("Weather K")
	w.Resize(fyne.Size{Width: 500, Height: 700})

	// Создание поля ввода
	entryToFind := widget.NewEntry()   // Создание поля ввода
	entryToFind.SetPlaceHolder("City") // placeholder "City"

	// Получение значения data из кнопки btn
	answare := widget.NewLabel("")

	// Кнопка Find и обработчик
	btn := widget.NewButton("Find", func() { //функция обработчик

		city := entryToFind.Text //достаем введенный текст

		// обрабатываем city
		data, err := config.Target(city)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(data)

	})

	w.SetContent(
		container.NewVBox(
			entryToFind,
			btn,
			answare,
		))

	w.ShowAndRun()
}

// Отправить запрос
// распарсить
// вывести на экран
