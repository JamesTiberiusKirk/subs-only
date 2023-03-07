FROM golang:alpine as builder
RUN mkdir /build
# ADD . /build
RUN mkdir /build/stripe-events-handler
ADD ./go.mod /build/
ADD ./go.sum /build/
ADD ./stripe-events-handler /build/stripe-events-handler

WORKDIR /build 
RUN go get ./... 
RUN go mod vendor
RUN mkdir /build/bin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin ./stripe-events-handler/./...

FROM alpine
COPY --from=builder /build/bin/stripe-events-handler /app/
WORKDIR /app
CMD ["./stripe-events-handler"]
