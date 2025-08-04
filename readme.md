FROM golang:1.24-bookworm AS build

RUN apt-get update && apt-get install -y --no-install-recommends \
 ca-certificates \
 git \
 && rm -rf /var/lib/apt/lists/\*

ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy.golang.crt
RUN update-ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /bin/app

FROM debian:bookworm-slim

ENV GRPCURL_VERSION=1.8.9

RUN apt-get update && apt-get install -y curl ca-certificates tar \
 && curl -L "https://github.com/fullstorydev/grpcurl/releases/download/v${GRPCURL_VERSION}/grpcurl_${GRPCURL_VERSION}_linux_x86_64.tar.gz" \
 -o grpcurl.tar.gz \
 && tar -xzf grpcurl.tar.gz \
 && mv grpcurl /usr/local/bin/grpcurl \
 && chmod +x /usr/local/bin/grpcurl \
 && rm -f grpcurl.tar.gz \
 && apt-get clean && rm -rf /var/lib/apt/lists/\*

COPY --from=build /bin/app /bin

EXPOSE 50051

CMD [ "/bin/app" ]
