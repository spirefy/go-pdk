package types

type Event struct {
	Id     string `json:"id"`
	Data   []byte `json:"data"`
	Source string `json:"source"`
	Target string `json:"target"`
}
