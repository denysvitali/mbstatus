build:
	go build -o build/mbstatus ./
build-mipsle:
	CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -o build/mbstatus-mipsle ./

.PHONY: build build-mips