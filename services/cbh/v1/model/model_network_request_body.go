package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 检查网络请求body
type NetworkRequestBody struct {

	// 状态
	Type string `json:"type"`

	// 安全组信息
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 网卡信息
	Nics []Nics `json:"nics"`
}

func (o NetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkRequestBody struct{}"
	}

	return strings.Join([]string{"NetworkRequestBody", string(data)}, " ")
}
