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
		{
			name: "Test Case 2 | Password length less than 6 characters",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "123",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 3,
			},
			wantErr: false,
		},
		{
			name: "Test Case 3 | Password length more than 20 characters",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "12$s4ode789G1*o0d1e2s3t4i5n6g7",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 3,
			},
			wantErr: false,
		},
		{
			name: "Test Case 4 | Password contains upper, lower, and number",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "76k%7Gzj!Q(X%s1g",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 0,
			},
			wantErr: false,
		},
		{
			name: "Test Case 5 | Password not contains upper, lower, and number",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "password",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 2,
			},
		},
		{
			name: "Test Case 6 | Password not contain 3 repeating characters",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "123FcL)#$456",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 0,
			},
			wantErr: false,
		},
		{
			name: "Test Case 7 | Password contain 3 repeating characters",
			args: args{
				passwordStep: &domain.StrongPasswordStepDtO{
					Password: "123FcccL)#$456",
				},
			},
			want: domain.StrongPasswordStepResponse{
				NumOfSteps: 1,
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
