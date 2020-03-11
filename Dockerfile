########################
### Builder          ###
########################
FROM golang:latest as builder

WORKDIR /go/src/github.com/asobti/kube-monkey
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

###
COPY . /go/src/github.com/asobti/kube-monkey
WORKDIR /go/src/github.com/asobti/kube-monkey
RUN make build

########################
### Final            ###
########################
FROM scratch
COPY --from=builder /go/src/github.com/asobti/kube-monkey/kube-monkey /go/bin/kube-monkey
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ENTRYPOINT ["/go/bin/kube-monkey"]
