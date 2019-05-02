FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

# RUN apt-get install pkg-config

# RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/pkgconfig/

# RUN apt install librdkafka-dev

# RUN ./configure

# RUN make

# RUN make install

#RUN brew install pkg-config

#RUN go get librdkafka.redist -Version 1.0.0

#RUN apt install librdkafka-dev

#RUN ./configure && make install

#RUN apt-get install -y pkg-config lxc-dev

# RUN git clone https://github.com/edenhill/librdkafka.git
# RUN cd librdkafka
#RUN ./configure --prefix /usr
#RUN make
#RUN make install

#RUN PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/x86_64-linux-gnu/pkgconfig



 RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig

# RUN export PKG_CONFIG_PATH

# RUN . ~/.bashrc

# RUN apt add librdkafka

# RUN Install librdkafka-dev

# RUN <PackageReference Include="librdkafka.redist" Version="1.0.0"

# RUN git clone https://github.com/confluentinc/cp-docker-images
# RUN cd cp-docker-images
#RUN git checkout 5.2.1-post

# RUN create --driver virtualbox --virtualbox-memory 6000 confluent
# RUN apt  install linuxbrew-wrapper 

# RUN RUN apk add --no-cache librdkafka


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