package message 
 
import ( 
	b "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/confluentinc/confluent-kafka-go/kafka" 

	"github.com/heaptracetechnology/microservice-kafka/result" 
)

type Subscribe struct {
	Data      Data   `json:"data"`
	Endpoint  string `json:"endpoint"`
	Id        string `json:"id"`
	IsTesting bool   `json:"istesting"`
	Topic     string `json:"topic"`
	GroupId   string `json:"group_id"`
}

type Data struct {
	Interval     int64 `json:"interval"`
	InitialDelay int64 `json:"initial_delay"`
}

type RequestPayload struct {
	Data     map[string]string `json:"data"`
	Endpoint string            `json:"endpoint"`
	Id       string            `json:"id"`
}
type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type Produce struct {
	Topic    string `json:"topic"`
	Message    string `json:"message"`
}

//Kafka service
func consume(responseWriter http.ResponseWriter, request *http.Request) {


	decoder := json.NewDecoder(request.Body)

	var listener Subscribe
	errr := decoder.Decode(&listener)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          listener.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	contentType := "application/json"
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(sub.Endpoint),
		cloudevents.WithStructuredEncoding(),
	)


	contentType := "application/json"
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(listener.Endpoint),
		cloudevents.WithStructuredEncoding(),
	)
	if err != nil {
		log.Printf("failed to create transport, %v", err)
		return
	}

	cloudClient, err := cloudevents.NewClient(t,
		cloudevents.WithTimeNow(),
	)


	c.SubscribeTopics([]string{listener.Topic, "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			source, err := url.Parse(sub.Endpoint)
			event := cloudevents.Event{
				Context: cloudevents.EventContextV01{
					EventID:     listener.Id,
					EventType:   "consume",
					Source:      cloudevents.URLRef{URL: *source},
					ContentType: &contentType,
				}.AsV01(),
				Data: ,msg.Value,
			}
			resp, err := cloudClient.Send(context.Background(), event)
			if err != nil {
				log.Printf("failed to send: %v", err)
			}


			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}


//Kafka service
func produce(responseWriter http.ResponseWriter, request *http.Request) {


	decoder := json.NewDecoder(request.Body)

	var produce Produce
	errr := decoder.Decode(&produce)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
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

	// Produce messages to topic (asynchronously)
		produceErr := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &produce.Topic, Partition: kafka.PartitionAny},
			Value:          []byte(produce.Message),
		}, nil)
		if produceErr != nil {
			log.Printf("failed to send: %v", err)
		}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}





//Kafka service
func CreateTopic(responseWriter http.ResponseWriter, request *http.Request) {


	decoder := json.NewDecoder(request.Body)

	var produce Produce
	errr := decoder.Decode(&produce)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	p, err := kafka.NewClient(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
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

	// Produce messages to topic (asynchronously)
		produceErr := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &produce.Topic, Partition: kafka.PartitionAny},
			Value:          []byte(produce.Message),
		}, nil)
		if produceErr != nil {
			log.Printf("failed to send: %v", err)
		}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

