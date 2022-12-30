FROM golang:1.20rc1-bullseye
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 8010
CMD ["app"]
