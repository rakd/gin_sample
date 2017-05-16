FROM alpine
#FROM gcr.io/google-appengine/golang

RUN apk update
RUN apk add ca-certificates
RUN apk --no-cache add curl

## Adjust time
#RUN ntpd -d -q -n -p 2.north-america.pool.ntp.org

## TIMEZONE
#RUN apk --update add tzdata && \
#    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
#    apk del tzdata
#
##RUN rm -rf /var/cache/apk/*

RUN mkdir /bin/app
COPY ./main /bin/main
COPY ./app/views /bin/app/views
COPY ./assets /bin/assets


EXPOSE 3000
WORKDIR /bin
ENTRYPOINT /bin/main
