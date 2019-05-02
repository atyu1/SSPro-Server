#Base image
FROM golang:1.9.2-alpine3.6 AS build

#Copy the code 
COPY main /root/
COPY templates /root/

# Default command
CMD ["/root/main"]
