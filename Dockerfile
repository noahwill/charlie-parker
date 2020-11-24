FROM golang:alpine as builderbase
RUN apk update && apk add gcc g++
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go test -mod=vendor ./internal/helpers/...

FROM builderbase as builder
WORKDIR /build
ARG app
RUN go build -mod=vendor -o app cmd/$app/*.go


FROM alpine AS buildfinal
RUN apk add --no-cache tzdata
RUN apk update && apk add ca-certificates && apk add bash
COPY --from=builder /build/app /
WORKDIR /
ENTRYPOINT ["/app"]