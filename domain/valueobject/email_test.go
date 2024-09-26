package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	for _, v := range []struct {
		testName string
		email    string
		errorMsg string
		hasError bool
	}{
		{email: "test@test.io", testName: "valid email"},
		{email: "testtest.io", testName: "Invalid email", hasError: true, errorMsg: "Invalid email"},
	} {
		t.Run(v.testName, func(t *testing.T) {
			email, err := NewEmail(v.email)
			if v.hasError {
				assert.NotNil(t, err)
				assert.Equal(t, v.errorMsg, err.Error())
				return
			}
			assert.Equal(t, v.email, email.email)
			assert.Nil(t, err)
		})
	}
}
