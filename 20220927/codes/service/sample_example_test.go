package service

import (
	"20220927/codes/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSampleService_Get_example(t *testing.T) {
	type field struct {
		SampleRepository func(ctrl *gomock.Controller) repository.IFSampleRepository
	}
	type arg struct {
		i int
	}
	type want struct {
		Value  string
		HasErr bool
	}
	tests := []struct {
		name  string
		field field
		arg   arg
		want  want
	}{
		{
			name: "正常系",
			field: field{
				SampleRepository: func(ctrl *gomock.Controller) repository.IFSampleRepository {
					mock := repository.NewMockIFSampleRepository(ctrl)
					mock.EXPECT().GetName(100).Return("Sample", nil)
					return mock
				},
			},
			arg: arg{
				i: 100,
			},
			want: want{
				Value:  "Sample",
				HasErr: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := SampleService{
				SampleRepository: tt.field.SampleRepository(ctrl),
			}
			got, err := s.Get(tt.arg.i)
			assert.Equal(t, tt.want.HasErr, err != nil)
			assert.Equal(t, tt.want.Value, got)
			ctrl.Finish()
		})
	}
}
