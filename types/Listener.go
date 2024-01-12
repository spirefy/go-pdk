package types

type Listener struct {
	Event    string `json:"event"`
	FuncName string `json:"funcName"`
}
