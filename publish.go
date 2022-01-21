package main

import (
	`encoding/json`
	`io/ioutil`
	`os`
	`path/filepath`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) publish(logger simaqian.Logger) (undo bool, err error) {
	packageFilepath := filepath.Join(p.config.Folder, packageJsonFilename)
	if _, statErr := os.Stat(packageFilepath); nil != statErr {
		undo = true
	}
	if undo {
		return
	}

	fields := gox.Fields{
		field.String(`filename`, packageFilepath),
		field.String(`registry`, p.config.Registry),
		field.Bool(`auth`, true),
		field.Bool(`ssl`, false),
	}
	pkg := new(_package)
	if data, readErr := ioutil.ReadFile(packageFilepath); nil != readErr {
		err = readErr
		logger.Warn(`读取文件出错`, fields.Connect(field.Error(err))...)
	} else {
		err = json.Unmarshal(data, pkg)
	}
	if nil != err {
		return
	}

	// 记录日志
	fields.Connects(gox.Fields{
		field.String(`name.short`, pkg.Name),
		field.String(`name.package`, pkg.packageName()),
		field.String(`name.full`, pkg.fullName(p.config.Username)),
		field.String(`version`, pkg.Version),
	})

	var exists bool
	if exists, err = p.exists(pkg.packageName()); nil != err {
		logger.Error(`检查仓库是否存在相同的包出错`, fields.Connect(field.Error(err))...)
	} else if exists {
		undo = true
		logger.Warn(`仓库中已经存在相同的包了`, fields...)
	}
	if nil != err || undo {
		return
	}

	logger.Info(`开始发布包到仓库`, fields...)

	args := []string{
		`publish`,
	}
	if `` != p.config.Access {
		args = append(args, `--access`, p.config.Access)
	}
	if `` != p.config.Tag {
		args = append(args, `--tag`, p.config.Tag)
	}
	// 发布包
	if err = p.npm(logger, args...); nil != err {
		logger.Error(`发布包到仓库出错`, fields.Connect(field.Error(err))...)
	}
	if nil != err {
		return
	}

	logger.Info(`发布包到仓库成功`, fields...)

	return
}
