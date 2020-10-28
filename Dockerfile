FROM iron/go:dev as build
LABEL maintainer="Galen Guyer <galen@galenguyer.com>"
WORKDIR /app
COPY main.go .
RUN go build -o haddock


FROM alpine:latest
WORKDIR /app
COPY --from=build haddock .
COPY words.txt .
COPY ./static/ ./static/
CMD ["/app/haddock"]
