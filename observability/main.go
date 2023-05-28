package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pai12345/Andromeda_Apps/observability/alert"
	"github.com/pai12345/Andromeda_Apps/observability/incident"
	"github.com/pai12345/Andromeda_Apps/observability/stream"
)

func main() {
	//alert
	alert_instance := alert.Alert{
		ID:          "ALR_ANDR12345",
		Kind:        "BusinessMetrics",
		Priority:    "P5",
		Severity:    "Warning",
		Description: "Order increase for Asus IPRO Laptops",
		Squad:       "Purchase Requisition",
		Tags:        map[string]string{"Kind": "BusinessMetrics", "type": "Alert", "Priority": "P5", "app": "ANDR"},
	}
	alert.CollectivePage(alert_instance)
	alert_instance.FetchAlert()

	//incident
	incident_instance := incident.Incident{
		ID:          "INC_ANDR12345",
		Alert:       &alert_instance,
		Kind:        "Service",
		Description: "Service Incident",
		Tags:        map[string]string{"Kind": "Service", "type": "Incident", "squad": "Purchase Requisition"},
	}
	incident_instance.FetchIncident()
	alert.CollectivePage(incident_instance)
	//stream
	stream_instance := stream.CStream{
		Source:      "https://kubernetes.default.svc",
		Destination: "https://ANDR.cls.in",
		Message:     "Hello World",
	}
	channel := make(chan string)
	go stream_instance.SendMessage("Hello world", channel)
	go stream_instance.SendMessage("Hallo Welt", channel)

	i := 0
	for result := range channel {
		fmt.Println(result)
		i++
		if i == 2 {
			close(channel)
		}
	}
	fmt.Println(stream_instance)

	//stream wait group
	stream_instance2 := stream.WStream{
		Data: stream.Keyvalue{
			"Source":      "https://kubernetes.default.svc",
			"Destination": "https://ANDR.cls.in",
			"Message":     "Hello World",
		},
	}

	var wg sync.WaitGroup
	for _, value := range stream_instance2.Data {
		wg.Add(1)
		go func(value string) {
			stream_instance2.SendMessage(value)
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // a job
		}(value)
	}
	wg.Wait()
}
