package model

import (
	"encoding/json"

	"strings"
)

type ShowInstanceUsersRespUsers struct {
	// 用户名称。

	UserName *string `json:"user_name,omitempty"`
	// 用户角色。

	Role *string `json:"role,omitempty"`
	// 是否为默认应用。

	DefaultApp *bool `json:"default_app,omitempty"`
	// 创建时间。

	CreatedTime *int64 `json:"created_time,omitempty"`
}

func (o ShowInstanceUsersRespUsers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersRespUsers struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersRespUsers", string(data)}, " ")
}
