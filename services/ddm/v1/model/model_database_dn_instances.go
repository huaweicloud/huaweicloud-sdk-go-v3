package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseDnInstances struct {

	// 逻辑库关联的DN实例的id
	Id string `json:"id"`

	// 关联DN实例的用户名。
	UserName string `json:"user_name"`

	// 关联DN实例的用户密码。
	UserPassword string `json:"user_password"`
}

func (o DatabaseDnInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseDnInstances struct{}"
	}

	return strings.Join([]string{"DatabaseDnInstances", string(data)}, " ")
}
