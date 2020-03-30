FROM golang:1.13

WORKDIR /go/src/covid-19-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["covid-19-api"]