package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnAccessPolicyResponseBodyAccessPolicy struct {

	// 访问策略ID
	Id *string `json:"id,omitempty"`
}

func (o CreateVpnAccessPolicyResponseBodyAccessPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnAccessPolicyResponseBodyAccessPolicy struct{}"
	}

	return strings.Join([]string{"CreateVpnAccessPolicyResponseBodyAccessPolicy", string(data)}, " ")
}
