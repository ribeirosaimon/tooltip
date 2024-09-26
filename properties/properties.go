package properties

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"gopkg.in/ini.v1"
)

var propertiesFile map[string][]byte

func NewPropertiesFile() {
	propertiesFile = make(map[string][]byte)

	args := os.Getenv(string(constants.AERGIA))
	properties := fmt.Sprintf("config.%s.properties", args)
	cfg, err := ini.Load(properties)
	if err != nil {
		panic("Fail to read properties file")
	}
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			re := regexp.MustCompile(`\{\{[A-Za-z0-9_]+\}\}`)
			if re.MatchString(key.Value()) {
				var environment string
				var replaceEnv string
				for _, u := range re.FindAllStringSubmatch(key.Value(), -1) {
					replaceEnv = u[0]
					environment = strings.ReplaceAll(u[0], "{{", "")
					environment = strings.ReplaceAll(environment, "}}", "")

				}
				getEnv := os.Getenv(environment)
				propertiesFile[key.Name()] = []byte(strings.ReplaceAll(key.Value(), fmt.Sprintf("%s", replaceEnv), getEnv))
			} else {
				logs.LOG.Message(fmt.Sprintf("Key: %s, Value: %s\n", key.Name(), key.Value()))
				propertiesFile[key.Name()] = []byte(key.Value())
			}
		}
	}
}

func NewMockPropertiesFile(mockedValues map[string][]byte) {
	propertiesFile = make(map[string][]byte)
	for key, value := range mockedValues {
		propertiesFile[key] = value
	}
}

func GetEnvironmentValue(v string) []byte {
	return propertiesFile[v]
}

func GetEnvironmentMode() constants.EnvironmentType {
	return constants.EnvironmentType(propertiesFile[string(constants.AERGIA)])
}
