package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationPolicyCreateReq 组织策略请求体
type OrganizationPolicyCreateReq struct {
	Policy *OrganizationPolicyCreate `json:"policy"`
}

func (o OrganizationPolicyCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyCreateReq struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyCreateReq", string(data)}, " ")
}
