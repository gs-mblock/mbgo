# demo1

来源：
[githun](https://github.com/donvito/learngo/tree/master/rest-kafka-mongo-microservice)

## run

$ brew install librdkafka  or  $ brew upgrade librdkafka

go run ./lib/mq/kafka/demo1/tokafka/rest-to-kafka.go
go run ./lib/mq/kafka/demo1/todata/kafka-to-mongo.go

## test

postman post ,post body-raw
[local](http://localhost:9090/jobs)

``` json
{
    "title":"title-s1",
    "description":"des1",
    "company":"xxgs",
    "salary":"12w"
}
```

TODO:
测试中，发现有问题，日后解决.19年初测试ok;
