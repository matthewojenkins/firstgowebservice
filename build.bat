
set GOOS = linux
set GOARCH = amd64

go build -o build\program .

docker build -f Docker\Dockerfile -t programservice:1.0 .

