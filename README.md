# Kafka Golang PoC

## Resources

- [Kafka Offset Explorer](https://kafkatool.com/download.html)
- [Apache Kafka with Golang - Getting Started](https://medium.com/swlh/apache-kafka-with-golang-227f9f2eb818)
- [sarama](https://pkg.go.dev/github.com/shopify/sarama#section-readme) -[Simple messaging framework using Go TCP server and Kafka](https://blog.gopheracademy.com/advent-2017/messaging-framework/)

Kafka Offset Explorer Configuration:

```
Kafka Cluster Version: 3.1
Zookeeper Host: localhost
Zookeeper Port: 2181
chroot path: /
Broker security type: Plaintext
Bootstrap servers: localhost:9092,localhost:9094,localhost:9096
```

## Init Kafka Cluster

```bash
docker compose up
```

## Init Producer

```bash
cd producer
go mod tidy
go run main.go
```

## Init Consumer

```bash
cd consumer
go mod tidy
go run main.go
```
