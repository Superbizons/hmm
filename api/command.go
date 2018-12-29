package api

type Command struct {
	Cmd string `json:"Command"`
}

type AuthorizationCommand struct {
	*Command
	Password string `json:"Password"`
}
