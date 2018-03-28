FROM alpine:latest

ADD cmd-server /cmd-server

EXPOSE 8080

ENTRYPOINT ["/cmd-server", "--host=0.0.0.0"]

