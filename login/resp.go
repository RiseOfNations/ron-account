package login

type Response struct {
	Token       string `json:"token"`
	Initialized bool   `json:"initialized"`
}

type Ok struct {
	Status string `json:"status"`
}
