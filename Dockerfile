FROM docker.yc345.tv/teacherschool/teacher-golang-run:latest

ADD ./dist/main /opt/app/
# ADD ./go-teacherschool-bus /go/bin/go-teacherschool-bus
ADD ./config /opt/app/config
#ADD ./doc /go/src/goal-management/doc
# ADD ./seelog.xml /go/src/main/seelog.xml

WORKDIR /opt/app/
# VOLUME /go/src/log
CMD ["/opt/app/main"]
# EXPOSE 3610