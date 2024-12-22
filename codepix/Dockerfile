FROM golang:1.22

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOPROXY=https://proxy.golang.org,direct

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    protobuf-compiler \
    librdkafka-dev \
    wget && \
    rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 && \
    go install github.com/spf13/cobra-cli@latest

RUN wget https://github.com/ktr0731/evans/releases/download/0.9.1/evans_linux_amd64.tar.gz && \
    tar -xzf evans_linux_amd64.tar.gz && \
    mv evans /usr/local/bin/ && \
    rm -f evans_linux_amd64.tar.gz

CMD ["tail", "-f", "/dev/null"]
