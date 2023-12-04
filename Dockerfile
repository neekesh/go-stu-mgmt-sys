FROM golang:1.19.3-bullseye

# Required because go requires gcc to build
# RUN apk add build-base

# RUN apk add inotify-tools

RUN echo $GOPATH

COPY . /SMS

WORKDIR /SMS

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go mod download

CMD sh /SMS/run.sh
