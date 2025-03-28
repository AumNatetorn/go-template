package template

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTemplateService_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMocktransactionRepo(ctrl)

	req := Request{
		ID: 1,
	}

	type args struct {
		req Request
	}

	argsMock := args{
		req: req,
	}

	testcases := []struct {
		name     string
		args     args
		mock     func()
		expected error
	}{
		{
			name: "case success",
			args: argsMock,
			mock: func() {
				mockRepo.EXPECT().Find(req).Return(&Response{
					Name: "test",
					Age:  10,
				}, nil)
			},
			expected: nil,
		},
	}

	s := NewTemplateService(mockRepo)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			_, err := s.Process(req)
			if tc.expected != nil {
				assert.EqualError(t, err, tc.expected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
