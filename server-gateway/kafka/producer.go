package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
)

func NewKafkaProducer(logger log.Logger) sarama.SyncProducer {
	brokers := []string{"172.20.10.7:9092"}

	kConf := sarama.NewConfig()
	kConf.Producer.Partitioner = sarama.NewRandomPartitioner
	kConf.Producer.RequiredAcks = sarama.WaitForAll
	kConf.Producer.Retry.Max = 5
	kConf.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, kConf)

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