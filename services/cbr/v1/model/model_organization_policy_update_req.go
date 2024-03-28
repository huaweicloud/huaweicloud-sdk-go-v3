package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationPolicyUpdateReq 更新组织策略请求体
type OrganizationPolicyUpdateReq struct {
	Policy *OrganizationPolicyUpdate `json:"policy"`
}

func (o OrganizationPolicyUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyUpdateReq struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyUpdateReq", string(data)}, " ")
}
