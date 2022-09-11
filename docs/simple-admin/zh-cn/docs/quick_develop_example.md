## 快速开发demo

## 安装goctls
- 方法1： 直接下载最新release
### ***建议: 使用 v0.0.1 , 目前官方的goctl修改后有bug, 无法生成go-swagger***
[Release](https://github.com/suyuan32/simple-admin-tools/releases)

将文件放到 $GOPATH/bin 目录下

- 方法2: 下载代码编译

```shell
git clone https://github.com/suyuan32/simple-admin-tools.git

cd tools/goctl

go build -o goctls goctl.go

cp ./goctls $GOPATH/bin/goctls
```

初始化环境

```shell
goctls env check -i -f --verbose
```

也可直接使用官方goctl，不过官方没有gorm支持和go-swagger支持

## RPC服务例子

### 首先修改 rpc/core.proto

```protobuf
syntax = "proto3";

package core;

option go_package = "./core";

// base message
message Empty {}

message IDReq {
  uint64 ID = 1;
}

message UUIDReq {
  string UUID = 1;
}

message BaseResp {
  string msg = 1;
}

message PageInfoReq {
  uint64 page = 1;
  uint64 page_size = 2;
}

// user message

message LoginReq {
  string username = 1;
  string password = 2;
}

message LoginResp {
  string id = 1;
  string role_name = 2;
  string role_value = 3;
  uint32 role_id = 4;
}

message ChangePasswordReq {
  string uuid = 1;
  string old_password = 2;
  string new_password = 3;
}

message CreateOrUpdateUserReq {
  uint64 id = 1;
  string avatar = 2;
  uint32 role_id = 3;
  string mobile = 4;
  string email = 5;
  int32  status = 6;
  string username = 7;
  string nickname = 8;
  string password = 9;
}

message UserInfoResp {
  uint64 id = 1;
  string avatar = 2;
  uint32 role_id = 3;
  string mobile = 4;
  string email = 5;
  int32  status = 6;
  string username = 7;
  string UUID = 8;
  string nickname = 9;
  string roleName = 10;
  int64 create_at = 11;
  int64 update_at = 12;
  string roleValue = 13;
}

message UserListResp {
  uint32 total = 1;
  repeated UserInfoResp data = 2;
}

message GetUserListReq {
  uint64 page = 1;
  uint64 page_size = 2;
  string username = 3;
  string nickname = 4;
  string email = 5;
  string mobile = 6;
  uint64 role_id = 7;
}

message UpdateProfileReq {
  string uuid = 1;
  string nickname = 2;
  string email = 3;
  string mobile = 4;
  string avatar = 5;
}

// menu messages

message CreateOrUpdateMenuReq {
  uint32 level = 1;
  uint32 parent_id = 2;
  string path = 3;
  string name = 4;
  string redirect = 5;
  string component = 6;
  uint32 order_no = 7;
  bool disabled = 8;
  Meta meta = 9;
  uint64 id = 10;
  uint32 menu_type = 11;
}

message MenuInfo {
  uint64 id = 1;
  int64 create_at = 2;
  int64 update_at = 3;
  uint32 level = 4;
  uint32 parent_id = 5;
  string path = 6;
  string name = 7;
  string redirect = 8;
  string component = 9;
  uint32 order_no = 10;
  bool disabled = 11;
  Meta meta = 12;
  repeated MenuInfo children = 13;
  uint32 menu_type = 14;
}

message Meta {
  bool keep_alive = 1;
  bool hide_menu = 2;
  bool hide_breadcrumb = 3;
  string current_active_menu = 4;
  string title = 5;
  string icon = 6;
  bool close_tab = 7;
}

message MenuInfoList {
  uint64 total = 1;
  repeated MenuInfo data = 2;
}

message MenuRoleInfo {
  uint64 id = 1;
  uint64 menu_id = 2;
  uint64 role_id = 3;
}

message MenuRoleListResp {
  uint64 total = 1;
  repeated MenuRoleInfo data = 2;
}

message CreateMenuParamReq {
  uint64 menu_id = 1;
  string type = 2;
  string key = 3;
  string value = 4;
}

message UpdateMenuParamReq {
  uint64 menu_id = 1;
  string type = 2;
  string key = 3;
  string value = 4;
  uint64 id = 5;
}

message MenuParamResp {
  uint64 id = 5;
  uint64 menu_id = 1;
  string type = 2;
  string key = 3;
  string value = 4;
  int64 create_at = 6;
  int64 update_at = 7;
}

message MenuParamListResp {
  uint64 total = 1;
  repeated MenuParamResp data = 2;
}

// role messages

message RoleInfo {
  uint64 id = 1;
  string name = 2;
  string value = 3;
  string default_router = 4;
  uint32 status = 5;
  string remark = 6;
  uint32 order_no = 7;
  int64 createAt = 8;
}

message RoleListResp {
  uint64 total = 1;
  repeated RoleInfo data = 2;
}

message SetStatusReq {
  uint64 id = 1;
  uint32 status = 2;
}

// casbin
message UpdatePolicyReq {
  uint64 role_id = 1;
  repeated PolicyPartInfo rules = 2;
}

message PolicyPartInfo {
  string path = 1;
  string method = 2;
}

message CreatePolicyReq {
  uint64 role_id = 1;
  PolicyPartInfo info = 2;
}

// api message
message ApiInfo {
  uint64 id = 1;
  int64 create_at = 2;
  string path = 3;
  string description = 5;
  string group = 6;
  string method = 7;
}

message ApiListResp {
  uint64 total = 1;
  repeated ApiInfo data = 2;
}

message ApiPageReq {
  PageInfoReq page = 1;
  string path = 2;
  string description = 3;
  string method = 4;
  string group = 5;
}

// authorization message

message RoleMenuAuthorityReq {
  uint64 role_id = 1;
  repeated uint64 menu_id = 2;
}
// return the role's authorization menu's ids
message RoleMenuAuthorityResp {
  repeated uint64 menu_id = 1;
}

// example
message HelloReq {
  string name = 1;
}

// service

service core {
  // init
  rpc initDatabase (Empty) returns (BaseResp);

  // user service
  rpc login(LoginReq) returns (LoginResp);
  rpc changePassword (ChangePasswordReq) returns (BaseResp);
  rpc createOrUpdateUser (CreateOrUpdateUserReq) returns (BaseResp);
  rpc getUserById (UUIDReq) returns (UserInfoResp);
  rpc getUserList (GetUserListReq) returns (UserListResp);
  rpc deleteUser (IDReq) returns (BaseResp);
  rpc updateProfile (UpdateProfileReq) returns (BaseResp);

  // menu service
  //menu management
  rpc createOrUpdateMenu (CreateOrUpdateMenuReq) returns (BaseResp);
  rpc deleteMenu (IDReq) returns (BaseResp);
  rpc getMenuListByRole (IDReq) returns (MenuInfoList);
  rpc getMenuByPage (PageInfoReq) returns (MenuInfoList);
  rpc createMenuParam (CreateMenuParamReq) returns (BaseResp);
  rpc updateMenuParam (UpdateMenuParamReq) returns (BaseResp);
  rpc deleteMenuParam (IDReq) returns (BaseResp);
  rpc getMenuParamById (IDReq) returns (MenuParamResp);
  rpc geMenuParamListById (IDReq) returns (MenuParamListResp);

  // role service
  rpc createOrUpdateRole (RoleInfo) returns (BaseResp);
  rpc deleteRole (IDReq) returns (BaseResp);
  rpc getRoleById (IDReq) returns (RoleInfo);
  rpc getRoleList (PageInfoReq) returns (RoleListResp);
  rpc setRoleStatus (SetStatusReq) returns (BaseResp);

  // api management service
  rpc createOrUpdateApi (ApiInfo) returns (BaseResp);
  rpc deleteApi (IDReq) returns (BaseResp);
  rpc getApiList (ApiPageReq) returns (ApiListResp);

  // authorization management service
  rpc getMenuAuthority (IDReq) returns (RoleMenuAuthorityResp);
  rpc createOrUpdateMenuAuthority (RoleMenuAuthorityReq) returns (BaseResp);

  // example
  rpc hello (HelloReq) returns (BaseResp);
}
```

添加 example rpc接口

在rpc目录下运行

```shell
goctls rpc protoc core.proto --proto_path=/home/ryan/GolandProjects/simple-admin-core/rpc/ --go_out=./types --go-grpc_out=./types --zrpc_out=./
```

proto_path需要绝对路径

修改 internal/logic/hellologic.go

```go
package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// example
func (l *HelloLogic) Hello(in *core.HelloReq) (*core.BaseResp, error) {
	return &core.BaseResp{Msg: in.Name}, nil
}

```

然后在 api/api_desc/ 目录下添加 example.api

```api
syntax = "v1"

info(
    title: "example"
    desc: "example"
    author: "ryansu"
    email: "yuansu.china.work@gmail.com"
    version: "v1.0"
)

type (
    // Hello response | Hello返回信息
    // swagger:response HelloResp
    HelloResp {
        // Msg
        Msg    string  `json:"msg"`

    }

    // swagger:parameters hello
    // Hello request | Hello请求
    HelloReq {
        // Name | 名称
        // Required: true
        Name   string `path:"name"`
    }
)

@server(
    group: example
)

service core {
    // swagger:route POST /example/hello/:name example hello
    // Hello | Hello
    // Responses:
    //   200: HelloResp
    //   401: HelloResp
    //   500: HelloResp
    @handler hello
    get /example/hello/:name (HelloReq) returns (HelloResp)
}

```

修改 core.api

```api
syntax = "v1"

info(
	title: "core service"
	desc: "this is the api discribe file for core services"
	author: "ryansu"
	email: "yuansu.china.work@gmail.com"
	version: "v1.0"
)

import "role.api"
import "user.api"
import "menu.api"
import "captcha.api"
import "apis.api"
import "authority.api"
import "example.api"         # here
import "base.api"

@server(
	group: core
)

service core {
	// swagger:route get /core/health core healthCheck
	// Check the system status | 检查系统状态
	@handler healthCheck
	get /core/health
	
	// swagger:route get /core/init/database core initDatabase
	// Initialize database | 初始化数据库
	// Responses:
	//   200: SimpleMsg
	//   500: SimpleMsg
	@handler initDatabase
	get /core/init/database returns (SimpleMsg)
}
```

在 api_desc目录下执行
```shell
goctls api go -api core.api -dir ..
```

修改 api/internal/logic/example/hellologic.go

```go
package example

import (
	"context"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloReq) (resp *types.HelloResp, err error) {
	result, err := l.svcCtx.CoreRpc.Hello(l.ctx, &core.HelloReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &types.HelloResp{Msg: result.Msg}, nil
}

```

### 启动 rpc 和 api

分别在 api rpc 目录下执行

```shell
go run core.go -f etc/core_dev.yaml 
```

### 浏览器访问

http://localhost:8500/example/hello/ryan

可以看到

{"msg":"ryan"}

如果添加了jwt则需要在角色权限页面配置权限
![pic](../../assets/add_example_api.png)
![pic](../../assets/add_example_api_authority.png)

## web修改

如果要在网页上调用
首先修改代码中的api

simple-admin-backend-ui/src/api/sys

添加 example.ts

```typescript
import { defHttp } from '/@/utils/http/axios';
import { ErrorMessageMode } from '/#/axios';
import { BaseIDReq, BasePageReq, BaseResp } from '/@/api/model/baseModel';
import { ApiInfo, ApiListResp } from './model/apiModel';

enum Api {
  Hello = '/sys-api/example/hello',
}

/**
 * @description: Get hello msg
 */

export const getApiList = (name: string) => {
  return defHttp.get<BaseResp>({ url: Api.GetApiList + '/' + name });
};
```

在 src/views/sys 中添加view即可