FROM golang:1.25.5 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN  go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o authentication ./app/cmd

FROM alpine
WORKDIR /app
COPY --from=builder ./app/authentication .
EXPOSE 8081
CMD [ "./authentication" ]