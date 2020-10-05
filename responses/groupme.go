package responses

// Message represents a GroupMe message
type Message struct {
	Message InnerMessage `json:"message"`
}

// InnerMessage is to get around GroupMe's dumb format
type InnerMessage struct {
	SourceGUID string `json:"source_guid"`
	Text       string `json:"text"`
}
