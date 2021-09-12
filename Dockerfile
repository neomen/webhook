FROM golang:1.17-alpine
LABEL maintainer "Yaroslav - https://github.com/neomen"
RUN apk --update add git bash
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /webhook
EXPOSE 8000
CMD [ "/webhook" ]