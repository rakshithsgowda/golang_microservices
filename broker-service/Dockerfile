FROM golang:1.25.5-alpine as builder
WORKDIR /app
COPY go.mod so.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0  GOOS=linux go build -o broker ./app/cmd

FROM alpine
WORKDIR /app
COPY --from=builder ./app/broker .
EXPOSE 8080
CMD [ "./broker" ]
