package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VerifyHostV2Req struct {
	// 主机端口

	Port int32 `json:"port"`
	// 主机账号

	Account string `json:"account"`
	// 主机密码

	Password string `json:"password"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
}

func (o VerifyHostV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyHostV2Req struct{}"
	}

	return strings.Join([]string{"VerifyHostV2Req", string(data)}, " ")
}
