package kafka

/*
import (
	"fmt"
	"github.com/Shopify/sarama"
	"gochat/config"
)
var consumer sarama.Consumer
func init(){
    c := config.GetConfig()
    logType := c.MsgChannelType.ChannelType
    if "kafka" == logType{
        configKafka := sarama.NewConfig()
        host := []string{c.MsgChannelType.KafkaHosts}
        client, _ := sarama.NewClient(host, configKafka)
        consumer, _ = sarama.NewConsumerFromClient(client)    
    }
    
}
func Receive(){
    partitionConsumer, _ := consumer.ConsumePartition("default_message", 0, sarama.OffsetNewest)
    defer partitionConsumer.Close()
    msg := <-partitionConsumer.Messages()
    fmt.Println("customer:",string(msg.Value))
    
}*/
