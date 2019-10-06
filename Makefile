# ------ VARIABLES ------------
DOCKERBUILDCMD=docker build
DOCKERRUNCMD=docker run
DOCKERNAME=atyu/sspro-server
GOMAINFILE=main
GOCMDLINUX=GOOS=linux go build -a -installsuffix cgo -o $(GOMAINFILE) .
GOCMDGET=go get
PORT=8080

TS=$(shell /bin/date +%s)

# ------- MAIN SECTION ---------
help:
	@echo "all    	    - build go for linux + docker container + run"
	@echo "install	    - build go fo linux + docker container"
	@echo "run          - start docker container"
	@echo "clean        - remove go binaries + containers"
	@echo "push         - git push commits to git"
	@echo "test-local   - test localy the code"

all: install run

install: build-linux build-container

#Build go package for linux platform
build-linux:
	$(GOCMDGET) "github.com/julienschmidt/httprouter"
	$(GOCMDLINUX)

#Create a docker container
build-container:
	$(DOCKERBUILDCMD) -t $(DOCKERNAME) .

#Run a container
run:
	$(DOCKERRUNCMD) --rm -it --name sspro-server -p $(PORT):$(PORT) -v /db/test.db:/db/test.db $(DOCKERNAME)

#ToDo: Add deploy 

#Clean binaries + conatiners
clean:
	rm -rf ./$(GOMAINFILE)	
	docker system prune -f

push:
	git push -u origin $(shell git rev-parse --abbrev-ref HEAD)

#------ TESTS ------
test-local: test-post test-get-all

test-post:
	curl -XPOST http://localhost:8080/datapoints -H "Content-Type: application/json" -d '{"data":[{"timestamp":$(TS), "location":"kosice", "room":"bedroom", "name":"test", "sensor":"temperature", "value":20}]}'

test-get-all:
	curl http://localhost:8080/datapoints/all/
