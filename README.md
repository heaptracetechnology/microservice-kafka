# Kafka as a microservice
An OMG service for kafka, it is a message queue service.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

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
$ omg subscribe topic consume -a topic=<TOPIC> -e HOST=<HOST> -e PORT=<PORT>
```
##### Kafka Produce
```sh
$ omg run produce -a topic=<TOPIC> -a message=<MESSAGE> -e HOST=<HOST> -e PORT=<PORT> 
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
