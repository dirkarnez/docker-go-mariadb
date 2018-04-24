FROM golang 

COPY again /go/bin

COPY ./app /go/src/eating.com/app
WORKDIR /go/src/eating.com/app
VOLUME ["/go/src/eating.com/app"]

RUN go get .
RUN go build
EXPOSE 5000
CMD again --bin=main run