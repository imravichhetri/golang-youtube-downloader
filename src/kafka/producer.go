package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	// "os/signal"
	// "syscall"
)

func Producer(videoUrl string) {

	broker := "localhost"
	topic := "youtube-dl"
	group := "group1"
	/*if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
			os.Args[0])
		os.Exit(1)
	}
	broker := os.Args[1]
	group := os.Args[3]
	topic := os.Args[2]*/
	// sigchan : =make( chan os.Signal, 1 );
	// signal.Notify( sigchan, syscall.SIGINT, syscall.SIGTERM )
	fmt.Println("Broker: " + broker + ", topics:" + topic + ", group: " + group)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Println("Failed to create producer: %\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer%v\n", p)

	deliveryChan := make(chan kafka.Event)

	// value := "Hello Go!"
	error := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(videoUrl), Key: []byte("Keys")}, deliveryChan)

	if error != nil {
		fmt.Println(error)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	fmt.Println(error, e, "error---")

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	// wait for delivery report goroutine to finish
	// _ = <-doneChan

	close(deliveryChan)

}
