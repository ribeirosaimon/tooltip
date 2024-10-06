package env

import (
	"fmt"
	"os"

	"github.com/ribeirosaimon/aergia-utils/tlog"
	"gopkg.in/yaml.v3"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type shortifyEnvironment struct {
	Config Config `yaml:"config"`
}

type Config struct {
	Env  Environment `yaml:"env"`
	Port int         `yaml:"port"`
}

var env shortifyEnvironment

func StartEnv(envName string) {
	f, err := os.Open(fmt.Sprintf("config.%s.yaml", envName))
	if err != nil {
		tlog.Warn("Failed to open config file", "err", err)
		// want to stop app, because dont have any environment
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&env); err != nil {
		tlog.Warn("Failed to parse config file", "err", err)
		panic(err)
	}
	tlog.Info(fmt.Sprintf("Loading environment variables in %s environment in port %d", env.Config.Env, env.Config.Port))
}

func GetConfig() Config {
	return env.Config
}
