FROM golang:1.20-alpine as builder

WORKDIR /src
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go main.go
COPY cmd/ cmd/
COPY internal/ internal/

RUN go build -o bin/server .

FROM scratch as package
WORKDIR /bin
COPY --from=builder /src/bin/server .
ENTRYPOINT [ "./server", "serve" ]