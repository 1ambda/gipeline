package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/go-kit/kit/log"
)

func NewKafkaProducer(logger log.Logger) sarama.SyncProducer {
	brokers := []string{"localhost:9092"}

	kConf := sarama.NewConfig()
	kConf.Producer.Retry.Max = 5
	kConf.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer(brokers, kConf)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	return producer
}

func DeleteKafkaProduce(logger log.Logger, p sarama.SyncProducer) {
	if err := p.Close(); err != nil {
		logger.Log("error", err)
	}
}