# Weather Metrics

This is a teeny-tiny web service that fetches the basic weather information for a given airport and serves it in the [Prometheus exposition format](https://prometheus.io/docs/instrumenting/exposition_formats/) at `/metrics`.

## Build

```bash
./script/build.sh
docker build -t <tagname> .
```

## Run

```bash
docker run --rm --env STATION=KDFW -p <hostport>:8080 <tagname>
```
