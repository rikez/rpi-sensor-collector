FROM raspbian/jessie


# Foundation
RUN DEBIAN_FRONTEND=noninteractive apt-get update -y && apt-get install --no-install-recommends -y -q pkg-config build-essential curl git gnupg apt-utils ca-certificates
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y -q -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew software-properties-common

# Golang Stuff
RUN curl -s https://dl.google.com/go/go1.11.10.linux-armv6l.tar.gz | tar -v -C /usr/local -xz
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

WORKDIR /go/src/github.com/rikez/rpi-sensor-collector
COPY . .
RUN go install -v ./...

CMD ["/go/bin/rpi-sensor-collector"]










