package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussUserForListAttribute 用户的权限属性。
type GaussDBforOpenGaussUserForListAttribute struct {

	// 用户是否具有超级用户权限，取值为“true”或“false”。
	Rolsuper *bool `json:"rolsuper,omitempty"`

	// 用户是否自动继承其所属角色的权限，取值为“true”或“false”。
	Rolinherit *bool `json:"rolinherit,omitempty"`

	// 用户是否支持创建其他子用户，取值为“true”或“false”。
	Rolcreaterole *bool `json:"rolcreaterole,omitempty"`

	// 用户是否可以创建数据库，取值为“true”或“false”。
	Rolcreatedb *bool `json:"rolcreatedb,omitempty"`

	// 用户是否可以登录数据库，取值为“true”或“false”。
	Rolcanlogin *bool `json:"rolcanlogin,omitempty"`

	// 用户连接实例的最大并发连接数。-1表示没有限制。
	Rolconnlimit *int32 `json:"rolconnlimit,omitempty"`

	// 用户是否属于复制角色，取值为“true”或“false”。
	Rolreplication *bool `json:"rolreplication,omitempty"`

	// 用户是否绕过每个行级安全策略，取值为“true”或“false”。
	Rolbypassrls *bool `json:"rolbypassrls,omitempty"`

	// 用户密码过期时间。
	Rolpassworddeadline *string `json:"rolpassworddeadline,omitempty"`
}

func (o GaussDBforOpenGaussUserForListAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserForListAttribute struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserForListAttribute", string(data)}, " ")
}
