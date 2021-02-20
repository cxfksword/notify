package wework

// payload struct holds message data
type payload struct {
	ToUser string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentID string `json:"agentid"`
	Text textMsg `json:"text"`
}

type textMsg struct {
   Content string `json:"content"`
}