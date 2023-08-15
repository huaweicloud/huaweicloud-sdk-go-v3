package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostedDirectConnectRequestBody 更新托管专线对象参数
type UpdateHostedDirectConnectRequestBody struct {
	HostedConnect *UpdateHostedDirectConnect `json:"hosted_connect,omitempty"`
}

func (o UpdateHostedDirectConnectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnectRequestBody", string(data)}, " ")
}
