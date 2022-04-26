# Alfred WorkHelper

## 介绍

本项目是一个简单的工具，用于在 Alfred 中快速查找工作相关的信息。

## 目前实现

整理来自日常工作中常用到的工具，其中包括：

- [x] 日期时间处理

- [x] 字符串常见编码

- [x] 字符串常见解码

- [x] 常见sign加密结果

- [x] 获取唯一ID

- [x] 获取内网所有可用的 IPV4 地址

- [ ] 模糊查询某个应用名进程id

- [ ] kill某个进程

## 插件直装

下载 [Alfred WorkHelper](https://github.com/yqchilde/alfred-workhelper/tree/main/alfred_workflow) 下的 `WorkHelper.alfredworkflow` 插件，双击可将其加入到 Alfred 中。

## 开发调试

1. 运行

```shell
# 记着修改脚本里面的参数来调试
bash workhelper.sh run 
```

2. 编译

```shell
bash workhelper.sh build
```

3. 打包

复制根目录生成的 `workHelper` 可执行文件到Alfred配置的Workflows中
