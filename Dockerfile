FROM --platform=$BUILDPLATFORM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -o ./webhook

FROM --platform=$BUILDPLATFORM alpine
RUN apk --update add git bash openssh-client
WORKDIR /app
COPY --from=builder /app/webhook /usr/bin/
ENTRYPOINT ["webhook"]
