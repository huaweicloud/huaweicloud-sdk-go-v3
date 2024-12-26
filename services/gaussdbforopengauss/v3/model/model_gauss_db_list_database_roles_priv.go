package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GaussDbListDatabaseRolesPriv struct {

	// 用户/角色是否具有管理员权限。
	Rolsuper *bool `json:"rolsuper,omitempty"`

	// 用户/角色是否自动继承其所属角色的权限。
	Rolinherit *bool `json:"rolinherit,omitempty"`

	// 用户/角色是否支持创建其他子用户。
	Rolcreaterole *bool `json:"rolcreaterole,omitempty"`

	// 用户/角色是否可以创建数据库。
	Rolcreatedb *bool `json:"rolcreatedb,omitempty"`

	// 用户/角色是否可以登录数据库。
	Rolcanlogin *bool `json:"rolcanlogin,omitempty"`

	// 用户/角色连接实例的最大并发连接数。-1表示没有限制。
	Rolconnlimit *int32 `json:"rolconnlimit,omitempty"`

	// 用户/角色是否属于复制角色。
	Rolreplication *bool `json:"rolreplication,omitempty"`

	// 用户/角色是否绕过每个行级安全策略。
	Rolbypassrls *bool `json:"rolbypassrls,omitempty"`

	// 用户/角色密码过期时间。
	Rolpassworddeadline *string `json:"rolpassworddeadline,omitempty"`
}

func (o GaussDbListDatabaseRolesPriv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDbListDatabaseRolesPriv struct{}"
	}

	return strings.Join([]string{"GaussDbListDatabaseRolesPriv", string(data)}, " ")
}
