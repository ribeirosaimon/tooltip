package properties

import (
	"fmt"
	"os"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"gopkg.in/ini.v1"
)

var propertiesFile map[string]string

func NewPropertiesFile() {
	propertiesFile = make(map[string]string)

	args := os.Getenv(string(constants.AERGIA))
	properties := fmt.Sprintf("config.%s.properties", args)
	cfg, err := ini.Load(properties)
	if err != nil {
		panic("Fail to read properties file")
	}
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			logs.LOG.Message(fmt.Sprintf("Key: %s, Value: %s\n", key.Name(), key.Value()))
			propertiesFile[key.Name()] = key.Value()
		}
	}
}

func NewMockPropertiesFile(mockedValues map[string]string) {
	propertiesFile = make(map[string]string)
	for key, value := range mockedValues {
		propertiesFile[key] = value
	}
}

func GetEnvironmentValue(v string) string {
	return propertiesFile[v]
}
