# ------ VARIABLES ------------
DOCKERBUILDCMD=docker build
DOCKERRUNCMD=docker run
DOCKERNAME=atyu/sspro-server
GOMAINFILE=main
GOCMDLINUX=CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(GOMAINFILE) .
GOCMDGET=go get
PORT=8080

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
	$(DOCKERRUNCMD) -p $(PORT):$(PORT) $(DOCKERNAME)

#ToDo: Add deploy 

#Clean binaries + conatiners
clean:
	rm -rf ./$(GOMAINFILE)	
	docker system prune -f

push:
	git push -u origin master

#------ TESTS ------
test-local: test-local-api

test-local-api:
	curl http://localhost:$(PORT)/test
	

