###
# BUILD STAGE
###
FROM golang:alpine AS build-stage

WORKDIR /psite
COPY . .
RUN go build main.go

###
# RUN STAGE
###
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /psite
COPY --from=build-stage /psite .

EXPOSE 3000
ENTRYPOINT /psite/main
