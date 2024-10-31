package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NacosUserInfo nacos用户信息。
type NacosUserInfo struct {

	// nacos用户名。
	UserName string `json:"user_name"`

	// nacos密码。
	Password string `json:"password"`
}

func (o NacosUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NacosUserInfo struct{}"
	}

	return strings.Join([]string{"NacosUserInfo", string(data)}, " ")
}
