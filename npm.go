package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) npm(logger simaqian.Logger, args ...string) (err error) {
	fields := gox.Fields{
		field.String(`exe`, npmExe),
		field.Strings(`args`, args...),
		field.Bool(`verbose`, p.config.Verbose),
		field.Bool(`debug`, p.config.Debug),
	}
	// 记录日志
	logger.Info(`开始执行Npm命令`, fields...)

	options := gex.NewOptions(gex.Args(args...), gex.Dir(p.config.Folder))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(npmExe, options...); nil != err {
		logger.Error(`执行Npm命令出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`执行Npm命令成功`, fields...)
	}

	return
}

func (p *plugin) exists(name string) (exists bool, err error) {
	output := ``
	if _, err = gex.Run(npmExe, gex.Args(`view`, name), gex.StringCollector(&output)); nil != err {
		return
	}
	exists = `` != output

	return
}
