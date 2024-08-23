package domain

// usecase
type StrongPasswordStepDtO struct {
	Password string `json:"password"`
	Step     int    `json:"step"`
}

type StrongPasswordStepResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
