
FROM golang:1.17.5 AS GO_BUILD
ENV CGO_ENABLED 0
COPY /api /api
WORKDIR /api
RUN go build -o server

FROM alpine:3.15
WORKDIR /api
COPY --from=GO_BUILD /api /api
EXPOSE 8080
CMD ["./server"]