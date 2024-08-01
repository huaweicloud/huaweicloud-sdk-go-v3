package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckNoNewAccessReqBody struct {

	// 该策略JSON格式策略文档。
	ExistingPolicyDocument string `json:"existing_policy_document"`

	// 该策略JSON格式策略文档。
	NewPolicyDocument string `json:"new_policy_document"`

	PolicyType *PolicyDocumentType `json:"policy_type"`
}

func (o CheckNoNewAccessReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNoNewAccessReqBody struct{}"
	}

	return strings.Join([]string{"CheckNoNewAccessReqBody", string(data)}, " ")
}
