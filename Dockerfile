FROM golang:1.27 AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o /out/demo .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/demo /demo
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/demo"]
