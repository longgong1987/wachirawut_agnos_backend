package controllers

type StrongPasswordStep struct {
	Password string `json:"password"`
	Step     int    `json:"step"`
}
