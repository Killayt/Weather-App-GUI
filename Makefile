SILENT: 

build: 
	go get fyne.io/fyne/v2
	go mod tidy
	go mod vendor
	go build

run: build
	./Weather-K-GUI
