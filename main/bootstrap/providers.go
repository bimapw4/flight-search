package bootstrap

type Providers struct {
}

type Error struct {
	Code    int    `json:"-"`
	Errors  any    `json:"errors,omitempty"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func Provider() Providers {

	return Providers{}
}
