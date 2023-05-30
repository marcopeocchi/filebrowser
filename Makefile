default:
	go build -o filebrowser main.go

vue:
	cd app && pnpm build

multiarch:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o build/filebrowser-armv6 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o build/filebrowser-armv7 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/filebrowser-arm64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/filebrowser-amd64 main.go