FROM golang:1.16.13

WORKDIR /newapp
COPY . .


RUN go build -o main APIS.go
CMD ["/newapp/main"]
