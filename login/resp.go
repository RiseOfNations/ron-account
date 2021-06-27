package login

type Response struct {
	Token string `json:"token"`
}

type Ok struct {
	Status string `json:"status"`
}
