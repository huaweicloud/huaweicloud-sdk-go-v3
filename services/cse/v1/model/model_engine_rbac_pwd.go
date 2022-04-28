package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作微服务引擎专享版安全认证信息
type EngineRbacPwd struct {

	// 开启安全认证的微服务引擎专享版默认root帐号的密码
	Pwd *string `json:"pwd,omitempty"`
}

func (o EngineRbacPwd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineRbacPwd struct{}"
	}

	return strings.Join([]string{"EngineRbacPwd", string(data)}, " ")
}
