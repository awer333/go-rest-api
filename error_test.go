package messagebird

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("Single", func(t *testing.T) {
		errRes := ErrorResponse{
			Errors: []Error{
				{
					Code:        42,
					Description: "something bad",
					Parameter:   "foo",
				},
			},
		}
		assert.Error(t, errRes)
	})

	t.Run("Multiple", func(t *testing.T) {
		errRes := ErrorResponse{
			Errors: []Error{
				{
					Code:        42,
					Description: "something bad",
					Parameter:   "foo",
				},
				{
					Code:        42,
					Description: "something else",
					Parameter:   "foo",
				},
			},
		}
		assert.Error(t, errRes)
	})
}

func TestErrorWithEmptyErrors(t *testing.T) {
	errRes := ErrorResponse{}
	assert.Error(t, errRes)
	assert.Equal(t, "API errors: ", errRes.Error())
}

func TestErrorWithOneEmptyError(t *testing.T) {
	errRes := ErrorResponse{Errors: []Error{
		{
			Code:        11,
			Description: "",
			Parameter:   "",
		},
	}}
	assert.Error(t, errRes)
	assert.Equal(t, "API errors: ", errRes.Error())
}
