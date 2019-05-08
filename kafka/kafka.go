package kafka 
 
import ( 
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/confluentinc/confluent-kafka-go/kafka" 
	"github.com/cloudevents/sdk-go"
	"github.com/heaptracetechnology/microservice-kafka/result" 
	"log"
	"net/url"
	"context"
	"time"
)

type Subscribe struct {
	Data      Data   `json:"data"`
	Endpoint  string `json:"endpoint"`
	Id        string `json:"id"`
	IsTesting bool   `json:"istesting"`
	GroupId   string `json:"group_id"`
}

type Data struct {
	Topic     string `json:"topic"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}


type Produce struct {
	Topic      string `json:"topic"`
	Message    string `json:"message"`
}

var Listener = make(map[string]Subscribe)
var rtmstarted bool
var isConsumerRunning bool

func Consume(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var listener Subscribe
	errr := decoder.Decode(&listener)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	
		c, _ := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "192.168.1.89",
			"group.id":          listener.Id,
			"auto.offset.reset": "earliest",
			"session.timeout.ms":    6000,
		})
	
	Listener[listener.Id] = listener
	if !rtmstarted {
		go KafkaRTM(*c)
		rtmstarted = true
	}

	bytes, _ := json.Marshal("Subscribed")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

func KafkaRTM(c kafka.Consumer) {
	istest := false
	quit := make(chan struct{})
	for {
		if len(Listener) > 0 {
			for k, v := range Listener {
				go getMessageUpdates(k, v,c)
				istest = v.IsTesting
			}
		} else {
			rtmstarted = false
			break
		}
		time.Sleep(3 * time.Second)
		if istest {
			close(quit)
			break
		}
	}
}


func getMessageUpdates(userid string, sub Subscribe,c kafka.Consumer) {

	contentType := "application/json"
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(sub.Endpoint),
		cloudevents.WithStructuredEncoding(),
	)

	if err != nil {
		log.Printf("failed to create transport, %v", err)
		return
	}

	cloudClient, err := cloudevents.NewClient(t,
		cloudevents.WithTimeNow(),
	)


	c.SubscribeTopics([]string{sub.Data.Topic, "^aRegex.*[Tt]opic"}, nil)
		msg, err := c.ReadMessage(-1)
		if err == nil {
			source, err := url.Parse(sub.Endpoint)
			event := cloudevents.Event{
				Context: cloudevents.EventContextV01{
					EventID:     sub.Id,
					EventType:   "consume",
					Source:      cloudevents.URLRef{URL: *source},
					ContentType: &contentType,
				}.AsV01(),
				Data:msg.Value,
			}
			resp, err := cloudClient.Send(context.Background(), event)
			if err != nil {
				log.Printf("failed to send: %v", err)
				fmt.Println(resp)
			}


			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
}



//Kafka service
func ProduceStream(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var produce Produce
	errr := decoder.Decode(&produce)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.1.89"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce produce)
	produceErr := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &produce.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(produce.Message),
	}, nil)
		if produceErr != nil {
			log.Printf("failed to send: %v", err)
			return
		} else {
			p.Flush(15 * 1000)

			message := Message{"true", "Topic created successfully.", http.StatusOK}
			bytes, _ := json.Marshal(message)
			result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
		}
		// Wait for message deliveries before shutting down
		
	
}





// //Kafka service
// func CreateTopic(responseWriter http.ResponseWriter, request *http.Request) {


// 	decoder := json.NewDecoder(request.Body)

// 	var produce Produce
// 	errr := decoder.Decode(&produce)
// 	if errr != nil {
// 		result.WriteErrorResponse(responseWriter, errr)
// 		return
// 	}

// 	p, err := kafka.NewClient(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer p.Close()

// 	// Delivery report handler for produced messages
// 	go func() {
// 		for e := range p.Events() {
// 			switch ev := e.(type) {
// 			case *kafka.Message:
// 				if ev.TopicPartition.Error != nil {
// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
// 				} else {
// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
// 				}
// 			}
// 		}
// 	}()

// 	// Produce messages to topic (asynchronously)
// 		produceErr := p.Produce(&kafka.Message{
// 			TopicPartition: kafka.TopicPartition{Topic: &produce.Topic, Partition: kafka.PartitionAny},
// 			Value:          []byte(produce.Message),
// 		}, nil)
// 		if produceErr != nil {
// 			log.Printf("failed to send: %v", err)
// 		}

// 	// Wait for message deliveries before shutting down
// 	p.Flush(15 * 1000)
// }

