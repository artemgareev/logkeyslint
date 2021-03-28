build:
	CGO_ENABLED=1 go build -buildmode=plugin plugin/logkeyslint.go