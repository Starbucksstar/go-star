FROM golang:latest AS build

WORKDIR /star

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gostar main.go

FROM alpine:latest

COPY --from=build /star/gostar .
ADD ./config.yaml .

EXPOSE 9090

CMD ["./gostar"]
