FROM golang:1.24-bullseye AS build

# RUN <<EOF
# groupadd --gid 1000 airbyte
# useradd --password "" --shell "/bin/false" --no-create-home --uid 1000 --gid 1000 airbyte
# mkdir /airbyte
# chown airbyte:airbyte /airbyte
# EOF

RUN mkdir /airbyte

WORKDIR /go/src/app
COPY . .

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM alpine:3.22

RUN <<EOF
addgroup -g 1000 airbyte
adduser -D -u 1000 -G airbyte airbyte
EOF

#FROM scratch
#COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=build /etc/passwd /etc/passwd
#COPY --from=build /etc/group /etc/group
COPY --from=build --chown=airbyte:airbyte /go/src/app/abcdk /
COPY --from=build --chown=airbyte:airbyte /airbyte /airbyte

ENV AIRBYTE_ENTRYPOINT="/abcdk"
USER airbyte:airbyte
ENTRYPOINT ["/abcdk"]