#Base image
FROM alpine

RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN mkdir /app

#Copy the code 
ADD main /app
ADD config.yaml /app

VOLUME /db

EXPOSE 8080

#Workdir is /usr/app
WORKDIR /app

# Default command
#ENTRYPOINT ["/app/main"]
CMD ["/app/main"]
