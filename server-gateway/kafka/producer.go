package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
	"strings"
	"time"
)

func NewKafkaProducer(logger log.Logger, brokers string) sarama.SyncProducer {
	logger.Log("brokers", brokers)
	brokerList := strings.Split(brokers, ",")

	kConf := sarama.NewConfig()
	kConf.Producer.Partitioner = sarama.NewRandomPartitioner
	kConf.Producer.RequiredAcks = sarama.WaitForAll
	kConf.Producer.Retry.Max = 5
	kConf.Producer.Return.Successes = true

	connected := false

	var err error = nil
	var producer sarama.SyncProducer

	for !connected {
		time.Sleep(5000 * time.Millisecond)
		producer, err = sarama.NewSyncProducer(brokerList, kConf)

		if err == nil {
			connected = true
		} else {
			logger.Log("message", "Failed to connect brokers. Retrying...")
		}
	}

	if err != nil {
		panic(err)
	}

	return producer
}

func DeleteKafkaProduce(logger log.Logger, p sarama.SyncProducer) {
	if err := p.Close(); err != nil {
		logger.Log("error", err)
	}
}