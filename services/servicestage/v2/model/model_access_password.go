package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessPassword struct {

	// 授权名称。
	Name string `json:"name"`

	// 仓库用户名。
	User string `json:"user"`

	// 仓库密码。
	Password string `json:"password"`
}

func (o AccessPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPassword struct{}"
	}

	return strings.Join([]string{"AccessPassword", string(data)}, " ")
}
