package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "order", 0)
	if err != nil {
		log.Println(err)
		return
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))
	batch := conn.ReadBatch(1e1, 1e9)

	data := make([]byte, 1e1)
	for {
		_, err := batch.Read(data)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(data))
	}
}
