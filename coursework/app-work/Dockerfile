FROM quay.io/projectquay/golang:1.25 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./main.go ./
COPY ./static ./static 

RUN CGO_ENABLED=0 GOOS=linux go build -o ./coursework
# multistage build - final stage
FROM alpine:3.22

EXPOSE 8080

WORKDIR /opt/app
COPY --from=builder /app .

CMD [ "./coursework" ]