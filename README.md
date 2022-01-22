# npm

Drone持续集成Npm插件，用于Npm打包发布

## 使用方式

```yaml
- name: 发布到Npm中央库
  image: dronestock/npm
  settings:
    folder: the library path
    username: xxx
    email: xxx@xxx.com
    password: xxx
    token: xxx
```
