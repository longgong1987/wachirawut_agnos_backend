package domain

// usecase
type StrongPasswordStep struct {
	Password string `json:"password"`
	Step     int    `json:"step"`
}
