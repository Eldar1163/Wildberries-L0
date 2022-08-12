package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"strings"
)

const (
	clusterID = "test-cluster"
	clientID  = "order-publisher"
	channel   = "order-notification"
)

func initConnection() stan.Conn {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		panic(err)
	}
	return sc
}

func readFilePathAndPublishCycle(sc stan.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter order json full filepath: ")
		text, _ := reader.ReadString('\r')
		text = strings.Trim(text, "\r\n\"")
		if text == "" {
			fmt.Println("Goodbye")
			return
		}

		file, err := os.ReadFile(text)
		if err != nil {
			fmt.Println("Cannot read a file, try another")
			continue
		}

		err = sc.Publish(channel, file)
		if err != nil {
			fmt.Println("Cannot publish message, please retry")
			continue
		}
	}
}

func main() {
	sc := initConnection()
	readFilePathAndPublishCycle(sc)
}
