package plugin

type Status struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewStatus(msg string, err string) *Status {
	return &Status{msg, err}
}
