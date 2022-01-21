package main

import (
	`github.com/dronestock/drone`
)

type plugin struct {
	config *config
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: &config{},
	}
}

func (p *plugin) Configuration() drone.Configuration {
	return p.config
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.login, drone.Name(`登录仓库`)),
		drone.NewStep(p.setup, drone.Name(`配置仓库`)),
		drone.NewStep(p.publish, drone.Name(`发布包`)),
	}
}
