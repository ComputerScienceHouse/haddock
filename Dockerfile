FROM golang:1.15.3-alpine3.12 as build
LABEL maintainer="Galen Guyer <galen@galenguyer.com>"
WORKDIR /app
COPY main.go .
RUN go build -o haddock


FROM alpine:latest
WORKDIR /app
COPY --from=build /app/haddock .
COPY words.txt .
COPY ./static/ ./static/
CMD ["/app/haddock"]
