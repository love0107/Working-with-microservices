package main

import (
	"context"
	"log"
	"log-service/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println("error writing mongo", err)
		return err
	}
	*resp = "Process Payload via RPC :" + payload.Name
	return nil
}
