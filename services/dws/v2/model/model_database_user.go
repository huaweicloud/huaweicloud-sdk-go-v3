package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户基本信息
type DatabaseUser struct {

	// 用户名
	Name *string `json:"name,omitempty"`

	// 是否可以登陆
	Login *bool `json:"login,omitempty"`
}

func (o DatabaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseUser struct{}"
	}

	return strings.Join([]string{"DatabaseUser", string(data)}, " ")
}
