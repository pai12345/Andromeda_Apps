package main

import (
	"github.com/pai12345/Andromeda_Apps/observability/alert"
	"github.com/pai12345/Andromeda_Apps/observability/incident"
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
}
