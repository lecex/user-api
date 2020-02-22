# user-api 用户

## MACOS开发环境
```
    # enable go modules
    export GO111MODULE=on

    go get github.com/micro/micro/v2

    # protobuf 开发环境
    go get -u github.com/golang/·/protoc-gen-go
    go get -u github.com/micro/protoc-gen-micro/v2

    go get -u github.com/gogo/protobuf/proto
    go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
    go get -u github.com/gogo/protobuf/gogoproto
```
## 依赖模块
- [user](https://github.com/lecex/user)
## Auth 认证服务
    - Auth              用户授权
    - Logout            用户推出
    - ValidateToken     授权认证
## Casbin 权限管理
    - AddPermission     增加权限
    - DeletePermissions 删除权限
    - UpdatePermissions 更新权限
    - GetPermissions    查询权限
    - AddRole           增加角色
    - DeleteRoles       删除角色
    - UpdateRoles       更新角色
    - GetRoles          查询角色
## FrontPermits 前端权限管理
    - All               全部权限
    - List              权限列表
    - Get               查询权限
    - Create            增加权限
    - Update            更新权限
    - Delete            删除权限
    - Sync              同步权限 (批量同步前端权限 需要最高 root 权限  *********高危 调用慎重*********)
## Health 健康检查
    - Health            健康检查
## Permissions 权限管理
    - All               全部权限
    - List              权限列表
    - Get               查询权限
    - Create            增加权限
    - Update            更新权限
    - Delete            删除权限··
## Roles 角色管理
    - All               全部角色
    - List              角色列表
    - Get               查询角色
    - Create            增加角色
    - Update            更新角色
    - Delete            删除角色
## Users 用户服务
    - Exist             查询用户
    - MobileBuild       绑定手机
    - SelfUpdate        更新登陆用户
    - Info              获取登陆用户信息
    - List              用户列表
    - Get               查询用户
    - Create            创建用户
    - Update            更新用户
    - Delete            删除用户