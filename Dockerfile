FROM golang:1.17.5

WORKDIR /newapp
COPY . .

RUN go build -o main APIS.go
CMD ["/newapp/main"]
