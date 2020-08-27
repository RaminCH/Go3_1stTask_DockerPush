FROM golang:latest

#-----------------------------
COPY go.mod .

COPY go.sum .

RUN go mod download 
#-----------------------------

#-----------------------------
RUN mkdir /app

ADD . /app


WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
#-----------------------------