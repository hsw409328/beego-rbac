#beego权限管理

#Rbac使用说明
        导入rbac.sql文件

        修改models/base.go
        
        orm.RegisterDataBase("default",
        "mysql","root:@tcp(127.0.0.1:3306)/rbac?charset=utf8",30)
        <br>
        改成自己的数据库连接地址

        运行main.go

        浏览器输入：http://localhost:8080<br>
        或者你服务器地址

        账号：admin 密码：admin
        该账号为超级权限。<br>
        与的配置在conf/app.conf superadmin对应

        新方法添加权限

        权限管理－添加权限

        角色管理－添加角色
        角色管理－配置权限

        Controller-Action里使用：
        _, action := this.GetControllerAndAction()
        this.CheckRbacAll(action)
