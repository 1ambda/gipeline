# gipeline

Example Data Pipeline Architecture written by Go with Kafka

![](http://benstopford.com/uploads/img/Slide41.png)

## Run
 
```
# kafka, server-gateway
$ docker-compose -f docker-compose.yml -f docker-compose-all.yml build
$ docker-compose -f docker-compose.yml -f docker-compose-all.yml rm 
$ docker-compose -f docker-compose.yml -f docker-compose-all.yml up 

# client
$ cd cli && go run main.go
```
 
## Development

**IMPORTANT**: Producer requires actual IP of brokers (we can't use `localhost` for kafka producers)

```
## kafka
KAFKA_ADVERTISED_HOST_NAME=192.168.0.35 docker-compose up   # should set advertised.host

## server-gateway
$ cd server-gateway
$ make install                             # only once
$ BROKERS=192.168.0.35:9092 make run-cont  # should set broker list

## cli
# cd cli
# go run main.go
```

## TODO

- server-connect
- elasticsearch
- web server
- wep application

## References

- Image: http://highscalability.com/blog/2015/5/4/elements-of-scale-composing-and-scaling-data-platforms.html