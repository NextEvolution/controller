package main

import (
	"github.com/nats-io/nats"
	"fmt"
	"nextevolution/common_types/controller_collector"
	"encoding/json"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to NATS: %s", err))
	}

	request := controller_collector.Request{
		FbToken: "some token",
		UserId: "1234",
		Groups: []string{"1601038443544679"},
		IgnoreAlbums: []string{},
		Keywords: []string{"sold"},
	}

	js, _ := json.Marshal(request)
	nc.Publish("collector", js)
	time.Sleep(10 * time.Second)
}