FROM golang 

COPY again /go/bin

COPY ./app /go/src/eating.com/app
WORKDIR /go/src/eating.com/app
VOLUME ["/go/src/eating.com/app"]

RUN go get github.com/go-sql-driver/mysql
RUN go build
CMD again --bin=main run 