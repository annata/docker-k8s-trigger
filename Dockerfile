FROM golang:alpine as builder
WORKDIR /k8s-trigger
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /k8s-trigger/k8s-trigger /
CMD ["/k8s-trigger"]