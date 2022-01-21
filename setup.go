package main

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) setup(logger simaqian.Logger) (undo bool, err error) {
	// 记录日志
	fields := gox.Fields{
		field.String(`registry`, p.config.Registry),
		field.Bool(`auth`, true),
		field.Bool(`ssl`, false),
	}
	logger.Info(`开始配置Npm仓库`, fields...)

	// 配置仓库地址
	if err = p.npm(logger, `config`, `set`, `registry`, p.config.Registry); nil != err {
		logger.Error(`配置Npm仓库失败`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	// 配置用户认证
	if err = p.npm(logger, `config`, `set`, `always-auth`, `true`); nil != err {
		logger.Error(`配置用户认证失败`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	// 配置安全连接
	if err = p.npm(logger, `config`, `set`, `strict-ssl`, `false`); nil != err {
		logger.Error(`配置不安全连接失败`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	logger.Info(`配置Npm仓库成功`, fields...)

	return
}
