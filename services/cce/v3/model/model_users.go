package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Users struct {
	// 当前为固定值“user“。

	Name *string `json:"name,omitempty"`

	User *User `json:"user,omitempty"`
}

func (o Users) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Users struct{}"
	}

	return strings.Join([]string{"Users", string(data)}, " ")
}
