release:
	GOOS=windows GOARCH=amd64 go build -o ./bin/azping_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/azping_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o ./bin/azping_darwin_amd64

push:
	setup/publish.sh
