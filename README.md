# Kafka as a microservice
An OMG service for kafka, it is a message queue service.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-kafka.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-kafka)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-kafka/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-kafka)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Kafka Consume
```sh
$ omg subscribe topic consume -a topic=<TOPIC> -e HOST=<HOST>
```
##### Kafka Produce
```sh
$ omg run produce -a topic=<TOPIC> -a message=<MESSAGE> -e HOST=<HOST>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-kafka .
```
### RUN
```
docker run -p 3000:3000 microservice-kafka
```
