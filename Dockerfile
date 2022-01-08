FROM golang:latest

WORKDIR /newapp
COPY . .


RUN go build -o main APIS.go
CMD ["/newapp/main"]
