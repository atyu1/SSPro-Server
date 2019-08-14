#Base image
FROM alpine:latest
MAINTAINER atyu atyu1@github.com

#Workdir is /usr/app
WORKDIR /usr/app

#Copy the code 
COPY main .
COPY config.yaml .

EXPOSE 8080

# Default command
CMD ["./main"]

