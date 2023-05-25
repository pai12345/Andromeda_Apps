package incident

import "github.com/pai12345/Andromeda_Apps/observability/alert"

type Incident struct {
	ID          string
	Alert       *alert.Alert
	Kind        string
	Description string
	Tags        map[string]string
}
