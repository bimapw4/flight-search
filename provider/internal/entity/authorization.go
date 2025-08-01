package entity

type Claim struct {
	UserID   int
	Username string
	IsAdmin  bool
	Exp      int
}

type Authorization struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
