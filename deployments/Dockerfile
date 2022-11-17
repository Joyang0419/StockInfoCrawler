# build stage
FROM golang:1.19-alpine3.16 AS build-env
ADD . /src
WORKDIR /src
RUN go mod download && go mod verify
RUN chmod -R 777 /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go
# final stage
FROM golang:1.19-alpine3.16
COPY --from=build-env /src/main /src/app.env /app/
WORKDIR /app
ENTRYPOINT ["./main"]
