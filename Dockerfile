FROM golang:1.19.1

WORKDIR /go/src

RUN useradd -ms /bin/bash golang && \
  echo "golang:golang" | chpasswd && adduser golang sudo

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=auto
ENV CGO_ENABLED=1

RUN apt-get update && \
  apt-get install build-essential protobuf-compiler librdkafka-dev sudo -y && \
  go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
  go get google.golang.org/protobuf/cmd/protoc-gen-go && \
  go get github.com/spf13/cobra && \
  wget https://github.com/ktr0731/evans/releases/download/v0.10.9/evans_linux_amd64.tar.gz && \
  tar -xzvf evans_linux_amd64.tar.gz && \
  mv evans ../bin && rm -rf evans_linux_amd64.tar.gz

RUN chown -R golang /go

USER golang

CMD ["tail", "-f", "/dev/null"]
