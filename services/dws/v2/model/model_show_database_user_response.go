package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDatabaseUserResponse struct {

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

	// 角色生效时间
	ValidBegin *int64 `json:"valid_begin,omitempty"`

	// 角色过期时间
	ValidUntil *int64 `json:"valid_until,omitempty"`

	// 是否锁定
	Lock           *bool `json:"lock,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseUserResponse", string(data)}, " ")
}
