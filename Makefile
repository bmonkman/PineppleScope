
build: 
	go build

container:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build -t pineapplescope-test -f Dockerfile .
