FROM golang:1.12

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV GIN_MODE release

RUN go build

RUN chmod +x /app/salamantex_back

ENTRYPOINT ["/app/salamantex_back"]
