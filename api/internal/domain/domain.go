package domain

type Usecase interface {
	CheckStrongPasswordStep(passwordStep *StrongPasswordStepDtO) (StrongPasswordStepResponse, error)
}

type Repositories interface {
	CreateStrongPasswordStep(strongPassword *StrongPasswordStepDtO) error
}

// usecase
type StrongPasswordStepDtO struct {
	Password string `json:"password"`
	Step     int    `json:"step"`
}

type StrongPasswordStepResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
