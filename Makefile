DOCKERBUILDCMD=sudo docker build
DOCKERRUNCMD=sudo docker run
DOCKERNAME=atyu/sspt-collector
GOMAINFILE=main
GOCMDLINUX=CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(GOMAINFILE) .
GOCMDGET=go get

PORT=8080

all: build-linux build-container

#Build go package for linux platform
build-linux:
	$(GOCMDGET) "github.com/julienschmidt/httprouter"
	$(GOCMDLINUX)

#Create a docker container
build-container:
	$(DOCKERBUILDCMD) -t $(DOCKERNAME) .

#Run a container
run:
	$(DOCKERRUNCMD) -p $(PORT):$(PORT) $(DOCKERNAME)

clean:
	rm -rf ./$(GOMAINFILE)	
