package main

import (
	`fmt`
	`net/url`

	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) login(logger simaqian.Logger) (undo bool, err error) {
	if `` != p.config.Token {
		err = p.loginWithToken(logger)
	} else if `` != p.config.Password {
		err = p.loginWithPassword(logger)
	}

	return
}

func (p *plugin) loginWithToken(logger simaqian.Logger) (err error) {
	var registry *url.URL
	if registry, err = url.Parse(p.config.Registry); nil != err {
		logger.Error(`配置的仓库不是有效的URL地址`, field.String(`registry`, p.config.Registry))
	}
	if nil != err {
		return
	}

	// 记录日志
	fields := gox.Fields{
		field.String(`username`, p.config.Username),
		field.String(`registry`, p.config.Registry),
	}
	logger.Info(`开始使用授权码方式登录仓库`, fields...)

	if err = p.npm(logger, `config`, `set`, fmt.Sprintf(`//%s:_authToken`, registry.Host), p.config.Token); nil != err {
		logger.Error(`使用授权码方式登录仓库失败`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`使用授权码方式登录仓库成功`, fields...)
	}

	return
}

func (p *plugin) loginWithPassword(logger simaqian.Logger) (err error) {
	// 记录日志
	fields := gox.Fields{
		field.String(`username`, p.config.Username),
		field.String(`registry`, p.config.Registry),
	}
	logger.Info(`开始使用密码方式登录仓库`, fields...)

	args := []string{
		`--username`,
		p.config.Username,
		`--password`,
		p.config.Password,
		`--email`,
		p.config.Email,
		`--registry`,
		p.config.Registry,
	}
	options := gex.NewOptions(gex.Args(args...))
	if p.config.Debug {
		options = append(options, gex.Quiet())
	}

	if _, err = gex.Run(addUserExe, options...); nil != err {
		logger.Error(`使用密码方式登录仓库失败`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`使用密码方式登录仓库成功`, fields...)
	}

	return
}
