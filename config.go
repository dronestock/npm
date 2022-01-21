package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}"`

	// 仓库地址
	Registry string `default:"${PLUGIN_REGISTRY=${REGISTRY=https://registry.npmjs.org/}}"`
	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${USERNAME}}" validate:"required"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${PASSWORD}}" validate:"required_without=Token"`
	// 授权码
	Token string `default:"${PLUGIN_TOKEN=${TOKEN}}" validate:"required_without=Password"`
	// 邮箱
	Email string `default:"${PLUGIN_EMAIL=${EMAIL}}" validate:"required_with=Password"`

	// 标签
	Tag string `default:"${PLUGIN_TAG=${TAG}}"`
	// 访问方式
	Access string `default:"${PLUGIN_ACCESS=${ACCESS=public}}" validate:"required,oneof=public private"`
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, c.Folder),

		field.String(`registry`, c.Registry),
		field.String(`username`, c.Username),
		field.String(`email`, c.Email),

		field.String(`tag`, c.Tag),
		field.String(`access`, c.Access),
	}
}
