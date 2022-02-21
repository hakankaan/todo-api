package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const filename = "../../configs/config.yaml"

type Service struct {
	CFG Config
}

// New is a constructor for config service
func New() (s Service) {
	c, err := getConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s = Service{
		CFG: *c,
	}

	return
}

// getConfig reads yaml and parse it to Config struct
func getConfig() (c *Config, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return
	}

	return
}
