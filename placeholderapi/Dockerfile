FROM golang:1.21 AS build-stage

WORKDIR /app

COPY . ./

RUN make deps

RUN go mod download

RUN make generate-api

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM build-stage AS run-test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

EXPOSE 8080

CMD ["/main"]