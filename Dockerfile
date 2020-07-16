FROM golang:1.14.5-alpine3.12 AS build

ENV GO111MODULE=on
RUN apk --no-cache add git make build-base
RUN apk --update-cache add tzdata
WORKDIR /go/src/github.com/Sw-Saturn/linebeacon-catchup
COPY . .

RUN mkdir -p /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o=/build/app cmd/main.go

FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Tokyo
COPY --from=build /build/app /build/app

ENV PORT=8080
ENTRYPOINT ["/build/app"]
