package incident

import "fmt"

func (inc Incident) FetchIncident() *Incident {
	fmt.Println(inc)
	return &inc
}

func (inc Incident) Page() bool {
	fmt.Printf("Squad %v paged for incident", inc.Alert.Squad)
	return true
}
