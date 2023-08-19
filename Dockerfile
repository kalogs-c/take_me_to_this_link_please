# Change this variable to the name of the app
ARG APP=app

ARG GO_VERSION=1.20.1

######################################

FROM golang:$GO_VERSION-alpine as builder

WORKDIR /$APP

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/main.go

######################################

FROM scratch

COPY --from=builder /$APP/$APP .

COPY --from=builder /$APP/templates /templates

EXPOSE 8080

ENTRYPOINT [ "./app" ]
