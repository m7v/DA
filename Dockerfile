FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/DA /opt/bin/DA
ADD ./views /opt/views
ADD ./static /opt/static

CMD ["/opt/bin/DA"]

