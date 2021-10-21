FROM golang:1.13 as build
WORKDIR /app
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN go build -v -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build ./app/ .
CMD ["./app"]
# ENTRYPOINT [ "sh" ]
# ENTRYPOINT ["sh", "go_jin_testing"]
