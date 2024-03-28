package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPoliciesRequest Request Object
type ListOrganizationPoliciesRequest struct {

	// 组织策略类型
	OperationType string `json:"operation_type"`
}

func (o ListOrganizationPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationPoliciesRequest", string(data)}, " ")
}
