package usecases

import (
	"testing"
	"wachirawut_agnos_backend/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateStrongPasswordStep(strongPassword *domain.StrongPasswordStepDtO) error {
	args := m.Called(strongPassword)
	return args.Error(0)
}

var (
	mockRepository   = &MockRepository{Mock: mock.Mock{}}
	quotationUsecase = Usecase{
		Repo: mockRepository,
	}
)

func TestCheckStrongPasswordStep(t *testing.T) {
	type args struct {
		passwordStep *domain.StrongPasswordStepDtO
	}
	tests := []struct {
		name    string
		args    args
		want    domain.StrongPasswordStepResponse
		wantErr bool
	}{
		{
			name: "Test Case 1 | Password length must be at least 6 characters and less than 20 characters",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "123456",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 2,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepository.On("CreateStrongPasswordStep", tt.args.passwordStep).Return(nil).Once()

			got, err := quotationUsecase.CheckStrongPasswordStep(tt.args.passwordStep)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckStrongPasswordStep error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckStrongPasswordStep = %v, want %v", got, tt.want)
			}
		})
	}

}
