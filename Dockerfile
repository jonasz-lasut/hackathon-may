FROM golang:1.20.3 as build

WORKDIR /build
COPY go.mod go.sum main.go Makefile ./
COPY server/ ./server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux make build


FROM scratch as production

COPY --from=build /build/main .

EXPOSE 5000
ENTRYPOINT [ "./main" ]
