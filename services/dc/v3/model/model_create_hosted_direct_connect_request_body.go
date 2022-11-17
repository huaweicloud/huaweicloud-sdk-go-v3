package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建托管专线对象信息
type CreateHostedDirectConnectRequestBody struct {
	HostedConnect *CreateHostedDirectConnect `json:"hosted_connect"`
}

func (o CreateHostedDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectRequestBody", string(data)}, " ")
}
