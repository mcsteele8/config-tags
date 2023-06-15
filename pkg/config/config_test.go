package config

import "testing"

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

func TestInitConfiguration(t *testing.T) {
	type args struct {
		cfg     interface{}
		verbose bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name:    "validate tags",
			wantErr: false,
			args: args{
				verbose: true,
				cfg:     cfg{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfiguration(&tt.args.cfg, tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("InitConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
