.SILENT:

build_windows: 
	go mod tidy
	go mod vendor
	go build -o Weather-K.exe

build:
	go mod tidy
	go mod vendor
	go build -o Weather-K

