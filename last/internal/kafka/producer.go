package kafka
import(
    "github.com/Shopify/sarama"
    "fmt"
    "gochat/config"
)
var (
    c = config.GetConfig().MsgChannelType
    producer sarama.AsyncProducer
    brokers = []string{c.KafkaHosts}
    topic   = c.KafkaTopic
)

func init(){
    logType := c.ChannelType
    if "kafka" == logType{
        config := sarama.NewConfig()
        config.Producer.Compression = sarama.CompressionGZIP
        client, _ := sarama.NewClient(brokers, config)
        fmt.Println("producer init success")
        producer, _ = sarama.NewAsyncProducerFromClient(client)    
    }
    
}
func SendLog(data string) {
	be := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: "default_message", Key: nil, Value: be}
}
