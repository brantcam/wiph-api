FROM golang:latest as builder

ENV GO111MODULES=on

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine:latest
COPY --from=builder /app .
EXPOSE 8080
CMD [ "./wiph-api" ]