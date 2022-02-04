FROM golang:alpine AS build

WORKDIR /TelegramBot
COPY go.* /TelegramBot//
RUN go mod download

COPY . /TelegramBot/
WORKDIR /TelegramBot/bot
RUN go build -o /go/bin/bot

FROM alpine AS run

COPY --from=build /go/bin/bot /
ENV GCLOUD_PROJECT=$GCLOUD_PROJECT
ENV TELEGRAM_APITOKEN=$TELEGRAM_APITOKEN

EXPOSE 8080
ENTRYPOINT ["./bot"]