FROM docker.io/library/golang:1.20-alpine3.17 as build
ADD . /app
WORKDIR /app
RUN go mod vendor && CGO=0 go build -ldflags="-s -w" -mod vendor -o ./bin/generate

FROM scratch
COPY --from=build /app/bin/generate /app/generate
EXPOSE 8080
ENTRYPOINT [ "/app/generate" ]