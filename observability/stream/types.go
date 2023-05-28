package stream

type Keyvalue map[string]string

type WStream struct {
	Data Keyvalue
}

type CStream struct {
	Source      string
	Destination string
	Message     string
}
