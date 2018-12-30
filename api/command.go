package api

type Command struct {
	Cmd string `json:"Command"`
}

type AuthorizationCommand struct {
	*Command
	Bots     int    `json"Bots"`
	Password string `json:"Password"`
}
