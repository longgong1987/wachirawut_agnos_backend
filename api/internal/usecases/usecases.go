package usecases

import (
	"unicode"
	"wachirawut_agnos_backend/internal/domain"
	"wachirawut_agnos_backend/internal/repositories"
)

type Usecase struct {
	Repo domain.Repositories
}

func NewUsecase(repo *repositories.Repositories) *Usecase {
	return &Usecase{
		Repo: repo,
	}
}

func (uc *Usecase) CheckStrongPasswordStep(passwordStep *domain.StrongPasswordStepDtO) (domain.StrongPasswordStepResponse, error) {

	var numOfSteps domain.StrongPasswordStepResponse

	passedPasswordLength := false
	containsUpperLowerNumber := false
	notContain3RepeatingCharacters := false
	hasUpper := false
	hasLower := false
	hasNumber := false

	// validate strong password step
	// password length must be at least 6 characters and less than 20 characters
	if len(passwordStep.Password) >= 6 && len(passwordStep.Password) < 20 {
		passedPasswordLength = true
	}

	// fmt.Println("passedPasswordLength: ", passedPasswordLength)

	for _, char := range passwordStep.Password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
		if unicode.IsNumber(char) {
			hasNumber = true
		}
	}

	if hasUpper && hasLower && hasNumber {
		containsUpperLowerNumber = true
	}

	// fmt.Println("containsUpperLowerNumber: ", containsUpperLowerNumber)

	// not contain 3 repeating characters
	foundRepeatingCharacters := false
	for i := 0; i < len(passwordStep.Password)-2; i++ {
		if passwordStep.Password[i] == passwordStep.Password[i+1] && passwordStep.Password[i] == passwordStep.Password[i+2] {
			foundRepeatingCharacters = true
		}
	}

	if !foundRepeatingCharacters {
		notContain3RepeatingCharacters = true
	}

	// fmt.Println("notContain3RepeatingCharacters: ", notContain3RepeatingCharacters)

	if passedPasswordLength && containsUpperLowerNumber && notContain3RepeatingCharacters {
		passwordStep.Step = 0
	} else if passedPasswordLength && containsUpperLowerNumber {
		passwordStep.Step = 1
	} else if passedPasswordLength {
		passwordStep.Step = 2
	} else {
		passwordStep.Step = 3
	}

	if err := uc.Repo.CreateStrongPasswordStep(passwordStep); err != nil {
		return numOfSteps, err
	}

	// set number of steps for response
	numOfSteps.NumOfSteps = passwordStep.Step

	return numOfSteps, nil
}
