package alert

type Alert struct {
	ID          string
	Kind        string
	Priority    string
	Severity    string
	Description string
	Squad       string
	Tags        map[string]string
}

type Levels map[string]bool

type Pager interface {
	Page() bool
}
