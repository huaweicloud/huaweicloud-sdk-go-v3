package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidatePolicyReqBody struct {

	// 该策略JSON格式策略文档。
	PolicyDocument string `json:"policy_document"`

	PolicyType *PolicyType `json:"policy_type"`

	ValidatePolicyResourceType *ValidatePolicyResourceType `json:"validate_policy_resource_type,omitempty"`
}

func (o ValidatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"ValidatePolicyReqBody", string(data)}, " ")
}
