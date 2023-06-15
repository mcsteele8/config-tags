package main

import "github.com/mcsteele8/config-tags/pkg/config"

var Cfg cfg

type cfg struct {
	Server      serverCfg
	LogLevel    string `env:"ORCHESTRATOR_LOG_LEVEL" default:"info"`
	Environment string `env:"ORCHESTRATOR_ENV" default:"sand"`
}

type serverCfg struct {
	Port        string `env:"ORCHESTRATOR_SERVER_PORT" default:"8081"`
	Deployments deploymentsCfg
}

type deploymentsCfg struct {
	K8sContext string `env:"ORCHESTRATOR_DEPLOYMENTS_K8S_CONTEXT" default:""`
}

func initConfig() {
	config.InitConfiguration(Cfg, true)
}

func main() {
	initConfig()
}
