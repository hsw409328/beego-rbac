# 基于beego2.x 权限管理

## <font color="green">升级说明</font>

```
1、升级beego2.x 版本支持
2、增加过滤器机制
3、新增配置不鉴权的接口和页面
4、新增配置只鉴权登录的接口和页面
5、修复同时登录并发写权限问题
5、修复页面部分使用和显示问题
```

# 初始化说明

## 1、导入rbac.sql文件

## 2、修改models/base.go

```go
//改成自己的数据库连接地址
orm.RegisterDataBase("default",
"mysql", "root:@tcp(127.0.0.1:3306)/rbac?charset=utf8", 30)
```

## 3、运行main.go

## 4、测试

```
浏览器输入：http://localhost:8080<br>
或者你服务器地址

账号：admin 密码：admin
该账号为超级权限。
```

# 配置文件支持

## 1、超级管理员

```
#修改为自己超级管理员账号，默认为admin
superadmin = admin
```

## 2、配置不鉴权的页面和路由

```
#默认前4个不允许修改，否则会重定向，在逗号后面增加自己编写的api或者页面地址
filterExcludeURL = /404,/exit,/login,/LoginSubmit,/test/*
```

## 3、配置只鉴权登录，不鉴权详细的页面和路由

```
#默认第一个不允许删除，否则系统会异常，使用方式同【2】部分
filterOnlyLoginCheckURL = /,/testCheckLogin/test
```

# 功能使用

```
使用admin、admin登录之后，会出现以下菜单
角色管理
    添加角色
    删除角色
    编辑角色
    配置角色与菜单的关系
权限管理（菜单/API）
    菜单/API列表
    编辑菜单
    删除菜单
    添加菜单
        参考已经内置到系统中的菜单即可
用户管理
    添加用户
    编辑用户
        添加和编辑的时候可以修改绑定的角色
```


