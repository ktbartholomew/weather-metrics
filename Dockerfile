FROM golang:1.7.5

ENV PORT=8080

EXPOSE 8080

COPY ./bin/weather-metrics /weather-metrics

ENTRYPOINT ["/weather-metrics"]
