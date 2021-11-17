VERSION = 0.1.$(shell date +%Y%m%d.%H%M)
FLAGS := "-s -w -X main.version=${VERSION}"


build:
	CGO_ENABLED=0 go build -ldflags=${FLAGS} -o lister .
	upx --lzma lister
	#CGO_ENABLED=0 go build -ldflags=${FLAGS} -o lister lister.go crypto.go
	#upx --lzma lister
	#CGO_ENABLED=0 go build -ldflags=${FLAGS} -o sender sender.go crypto.go
	#upx --lzma sender
