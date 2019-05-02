FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go


RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig


RUN curl -O http://packages.confluent.io/archive/5.2/confluent-5.2.1-2.12.zip

RUN curl -O http://packages.confluent.io/archive/5.2/confluent-5.2.1-2.12.tar.gz


RUN curl -O http://packages.confluent.io/archive/5.2/confluent-community-5.2.1-2.12.zip


RUN curl -O http://packages.confluent.io/archive/5.2/confluent-community-5.2.1-2.12.tar.gz


WORKDIR /go/src/github.com/heaptracetechnology/microservice-kafka

ADD . /go/src/github.com/heaptracetechnology/microservice-kafka

RUN go get github.com/confluentinc/confluent-kafka-go/kafka

RUN go install github.com/heaptracetechnology/microservice-kafka

ENTRYPOINT microservice-kafka

EXPOSE 3000