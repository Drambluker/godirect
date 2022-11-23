FROM golang:1.17

WORKDIR /usr/src/godirect

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/godirect main.go

CMD ["godirect", "/opt/godirect/conf/config.json"]
