package coparse

import (
	"os"
	"regexp"
	"strings"
)

var delimRegex *regexp.Regexp

func init() {
	delimRegex = regexp.MustCompile(`[:= ]`)
}

type Config map[string]interface{}

func (c *Config) Get(key string) any {
	_c := *c
	return _c[key]
}

func ParseConfig(path string) (config *Config, err error) {
	var configL = Config{}

	_, err = os.Stat(path)
	if err != nil {
		return
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return
	}

	flags := strings.Split(string(fileData), "\n")
	for _, flag := range flags {
		if len(flag) == 0 {
			continue
		}

		if delimRegex.MatchString(flag) {
			posArg := delimRegex.Split(flag, -1)
			configL[posArg[0]] = posArg[1]
		} else {
			configL[flag] = true
		}

		config = &configL
	}

	return
}

func MergeConfig(configs ...map[string]interface{}) {}

func FlagsToConfig(flags map[string]interface{}) {}
