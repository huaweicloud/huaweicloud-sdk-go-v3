package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新端口请求体
type UpdatePortOption struct {
	// IP/Mac对列表

	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`
	// 安全组列表

	SecurityGroups *[]string `json:"security_groups,omitempty"`
}

func (o UpdatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortOption struct{}"
	}

	return strings.Join([]string{"UpdatePortOption", string(data)}, " ")
}
