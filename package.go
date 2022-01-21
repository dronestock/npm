package main

import (
	`fmt`
	`path/filepath`
)

type _package struct {
	// 包名
	Name string `json:"name"`
	// 版本
	Version string `json:"version"`
	// 发布配置
	PublishConfig publishConfig `json:"publishConfig"`
}

func (p *_package) packageName() string {
	return fmt.Sprintf(`%s@%s`, p.Name, p.Version)
}

func (p *_package) fullName(username string) string {
	return filepath.Clean(fmt.Sprintf(`%s/%s/%s`, p.PublishConfig.Registry, username, p.Name))
}
