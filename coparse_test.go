package coparse_test

import (
	"testing"

	coparse "github.com/barelyhuman/coparse"
)

func TestParseConfig(t *testing.T) {
	config, err := coparse.ParseConfig("./config.test")
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		val := config.Get("bool")
		if val != true {
			t.Fail()
		}

		val = config.Get("pos")
		if val != "value" {
			t.Fail()
		}
	}

}
