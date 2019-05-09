FROM golang

RUN apt-get update 

RUN apt-get install git

RUN git clone https://github.com/edenhill/librdkafka.git \
    && cd librdkafka \
    && ./configure --prefix /usr \
    && make \
    && make install

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

RUN go get github.com/confluentinc/confluent-kafka-go/kafka

WORKDIR /go/src/github.com/heaptracetechnology/microservice-kafka

ADD . /go/src/github.com/heaptracetechnology/microservice-kafka

RUN go install github.com/heaptracetechnology/microservice-kafka

ENTRYPOINT microservice-kafka

EXPOSE 3000
