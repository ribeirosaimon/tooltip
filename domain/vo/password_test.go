package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassoword(t *testing.T) {
	for _, v := range []struct {
		testName string
		password string
		errorMsg string
		hasError bool
	}{
		{
			password: "P@sw0rd!",
			testName: "valid password",
		},
		{
			password: "Pasw0rdd",
			testName: "Invalid password because not have Special letter",
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
			hasError: true,
		},
		{
			password: "p@sw0rd!",
			testName: "Invalid password because not have Upercase letter",
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
			hasError: true,
		},
		{
			password: "p@sword!",
			testName: "Invalid password because not have number",
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
			hasError: true,
		},
		{
			password: "p@a!",
			testName: "Invalid password because too short length",
			errorMsg: "password too short",
			hasError: true,
		},
		{
			password: "",
			testName: "Invalid password because not have password",
			errorMsg: "password required",
			hasError: true,
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			password, err := NewPassword(v.password)
			if v.hasError {
				assert.NotNil(t, err)
				assert.Equal(t, v.errorMsg, err.Error())
				return
			}
			assert.Equal(t, v.password, password.GetValue())
			assert.Nil(t, err)
		})
	}
}
