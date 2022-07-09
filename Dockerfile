FROM alpine:3.16

# make sure it support go binary library
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# we don't compile golang at alpine, but we will compile it at somewhere else.
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

#APP HOME
ENV APP_HOME /biukop/web
RUN mkdir -p "$APP_HOME"

#update static html files
RUN mkdir -p $APP_HOME/html
COPY ./deploy/biukopweb-html $APP_HOME/html

#copy production configuration file
COPY ./deploy/config_production.json $APP_HOME/config.json
COPY ./goweb $APP_HOME/goweb

WORKDIR "$APP_HOME"
EXPOSE 8080

ENV PATH "$APP_HOME:$PATH"
CMD ["goweb", "-f", "config.json"]