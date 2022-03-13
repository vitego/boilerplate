build = ".build"

# Compile binary
compile:
	@mkdir -p .build
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./${build}/http -v ./cmd/http

# Run server
run/http:
	@make compile
	@./${build}/http