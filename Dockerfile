FROM golang:1.7.5

ENV PORT=8080
EXPOSE 8080
WORKDIR /src

COPY . /src
RUN CGO_ENABLED=0 go build -o /weather-metrics -a -installsuffix cgo

ENTRYPOINT ["/weather-metrics"]
