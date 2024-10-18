package tserver

import (
	"fmt"
	"os"

	"github.com/ribeirosaimon/tooltip/tlog"
	"gopkg.in/yaml.v3"
)

// Environment is a env value
type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type config struct {
	Env      Environment `yaml:"env"`
	Port     int         `yaml:"port"`
	HostName string      `yaml:"hostname"`
}

type shortifyEnvironment struct {
	Config config   `yaml:"config"`
	Mongo  dbConfig `yaml:"mongo"`
	Pgsql  dbConfig `yaml:"pgsql"`
	Redis  dbConfig `yaml:"redis"`
}

type dbConfig struct {
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	EntryPoint string `yaml:"entryPoint"`
}

var env shortifyEnvironment

func StartEnv(envName Environment) {
	f, err := os.Open(fmt.Sprintf("config.%s.yaml", envName))
	if err != nil {
		tlog.Warn("StartEnv", "Failed to open config file", "err", err)
		// want to stop app, because dont have any environment
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&env); err != nil {
		tlog.Warn("StartEnv", "Failed to parse config file", "err", err)
		panic(err)
	}
	tlog.Info("StartEnv", fmt.Sprintf("Loading environment variables in %s environment in port %d", env.Config.Env, env.Config.Port))
}

func GetEnvironment() config {
	return env.Config
}
func GetMongoConfig() dbConfig {
	return env.Mongo
}
func GetPgsqlConfig() dbConfig {
	return env.Pgsql
}

func GetRedisConfig() dbConfig {
	return env.Redis
}
