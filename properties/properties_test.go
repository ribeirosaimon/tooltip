package properties

import (
	"os"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestProperties(t *testing.T) {
	if err := os.Setenv(string(constants.AERGIA), "test"); err != nil {
		panic(err)
	}

	t.Run("open config properties", func(t *testing.T) {
		NewPropertiesFile()
	})

	value := GetEnvironmentValue("local.env")
	assert.Equal(t, "http://localhost", value)
}
