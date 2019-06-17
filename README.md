# _Kafka_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-kafka.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-kafka)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-kafka/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-kafka)

An OMG service for kafka, it is a message queue service.

## Direct usage in [Storyscript](https://storyscript.io/):

```coffee
>>> kafka produce topic:'topic' message:'messageForTopic'
{"success":"true/false","message":"success/failure message","statusCode":"HTTPstatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨


## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Kafka Produce
```shell
$ omg run produce -a topic=<TOPIC> -a message=<MESSAGE> -e HOST=<HOST>
```
##### Kafka Consume
```shell
$ omg subscribe topic consume -a topic=<TOPIC> -e HOST=<HOST>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/kafka/blob/master/LICENSE).
