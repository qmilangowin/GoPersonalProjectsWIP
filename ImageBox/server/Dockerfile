FROM golang:alpine AS build

WORKDIR /server
COPY go.* /server/
RUN go mod download

COPY . /server
WORKDIR /server/cmd
RUN go build -o /go/bin/server

FROM alpine AS run

COPY --from=build /go/bin/server /
ENV GCLOUD_PROJECT=$GCLOUD_PROJECT
ENV BUCKET=$BUCKET

EXPOSE 8080
ENTRYPOINT ["./server"]






