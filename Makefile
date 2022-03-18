build : clean amd64 arm64
clean :
	rm -rf target; go clean

amd64 :
	go clean; env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags=-static" -o target/linux-amd64/ddns ./cmd/ddns


arm64 :
	go clean; env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -extldflags=-static" -o target/linux-arm64/ddns ./cmd/ddns

