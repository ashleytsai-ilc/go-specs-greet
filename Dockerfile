FROM golang:1.20-alpine

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/*.go

EXPOSE 8080

CMD [ "./svr" ]