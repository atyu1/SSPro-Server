#Base image
FROM scratch

#Workdir is /usr/app
WORKDIR /usr/app

#Copy the code 
COPY main .
COPY templates templates

# Default command
CMD ["/usr/app/main"]

