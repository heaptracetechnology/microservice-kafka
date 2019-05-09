FROM golang

RUN apt-get update 
RUN apt-get upgrade -y
RUN apt-get install git

RUN git clone https://github.com/edenhill/librdkafka.git \
    && cd librdkafka \
    && ./configure --prefix /usr \
    && make \
    && make install    

RUN apt-get install pkg-config

RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/pkgconfig/

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

RUN go get github.com/confluentinc/confluent-kafka-go/kafka

WORKDIR /go/src/github.com/heaptracetechnology/microservice-kafka

ADD . /go/src/github.com/heaptracetechnology/microservice-kafka

RUN go install github.com/heaptracetechnology/microservice-kafka

ENTRYPOINT microservice-kafka

EXPOSE 3000