FROM golang:1.11-alpine as builder

# Install dependenices.
RUN apk update \
    && apk add --no-cache git make gcc autoconf libtool automake python3 musl-dev linux-headers

# Set Working Directory to Go folder.
WORKDIR /root/go/src/github.com/lightningnetwork

# Clone the LND project.
RUN git clone https://github.com/lightningnetwork/lnd

# Set Working Directory to LND.
WORKDIR ./lnd

# Install LND.
RUN make \ 
    && make install

# Download and Install BTCD.
WORKDIR /root/go/src/github.com
RUN git clone https://github.com/btcsuite/btcd.git

# Install btcd.
WORKDIR /root/go/src/github.com/btcd
RUN go install . ./cmd/...

FROM alpine:latest

COPY --from=builder /go/bin/btcctl /bin/btcctl
COPY --from=builder /go/bin/btcd /bin/btcd
COPY --from=builder /go/bin/lnd /bin/lnd
COPY --from=builder /go/bin/lncli /bin/lncli

# Set the entry workdir to simnet.
WORKDIR /simnet

# Expose ports for Bitcoin and LND.
EXPOSE 8332 8333 10009 8080 9735
