package api

type Command struct {
	Cmd string `json:"Command"`
}

type AuthorizationCommand struct {
	*Command
	Bots     int    `json:"Bots"`
	Port     int    `json:"Port"`
	Password string `json:"Password"`
}
