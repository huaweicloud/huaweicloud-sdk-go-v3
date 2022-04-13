package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessToken struct {
	// 授权名称。

	Name string `json:"name"`
	// git仓库设置中创建的私有token。

	Token string `json:"token"`
	// git仓库的主机地址，如https://192.168.1.1:8080/gitlab，默认为官方主机。

	Host *string `json:"host,omitempty"`
}

func (o AccessToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessToken struct{}"
	}

	return strings.Join([]string{"AccessToken", string(data)}, " ")
}
