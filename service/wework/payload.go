package wework

// payload struct holds message data
type payload struct {
	ToUser   string   `json:"touser"`
	MsgType  string   `json:"msgtype"`
	AgentID  string   `json:"agentid"`
	Textcard textCard `json:"textcard"`
}

type textCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
