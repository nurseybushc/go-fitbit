package fitbit

type BodyService service

type Body struct {
	Weight Weight `json:"weight"`
}

type Weight struct {
	Date   string `json:"date"`
	LogId  int    `json:"logId"`
	Time   string `json:"time"`
	Weight int    `json:"weight"`
	Source string `json:"source"`
}
