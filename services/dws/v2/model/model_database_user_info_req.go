package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseUserInfoReq 用户详细信息
type DatabaseUserInfoReq struct {

	// 是否可以登陆
	Login *bool `json:"login,omitempty"`

	// 创建角色权限
	Createrole *bool `json:"createrole,omitempty"`

	// 创建数据库权限
	Createdb *bool `json:"createdb,omitempty"`

	// 系统管理员
	Systemadmin *bool `json:"systemadmin,omitempty"`

	// 审计管理员
	Auditadmin *bool `json:"auditadmin,omitempty"`

	// 继承所在组权限
	Inherit *bool `json:"inherit,omitempty"`

	// 访问外表权限
	Useft *bool `json:"useft,omitempty"`

	// 连接数限制
	ConnLimit *int32 `json:"conn_limit,omitempty"`

	// 是否允许流复制
	Replication *bool `json:"replication,omitempty"`

	// 角色生效时间 yyyy-MM-ddTHH:mm:ssZ
	ValidBegin *string `json:"valid_begin,omitempty"`

	// 角色过期时间 yyyy-MM-ddTHH:mm:ssZ
	ValidUntil *string `json:"valid_until,omitempty"`

	// 是否锁定
	Lock *bool `json:"lock,omitempty"`
}

func (o DatabaseUserInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseUserInfoReq struct{}"
	}

	return strings.Join([]string{"DatabaseUserInfoReq", string(data)}, " ")
}
