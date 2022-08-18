# Dockerfile.production
FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_USER app
ENV APP_HOME /go/src/lolesports

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY . .

RUN ls -la /go/src/lolesports
RUN go mod download
RUN go mod verify
RUN go build -o lolesports /go/src/lolesports/server.go

FROM debian:buster
FROM registry.semaphoreci.com/golang:1.18

ENV APP_USER app
ENV APP_HOME /go/src/lolesports

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/lolesports $APP_HOME

EXPOSE 8083
USER $APP_USER

CMD ["./lolesports"]