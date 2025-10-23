package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties 保护类型：备份（backup）、复制(replication)。
type CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties struct {

	// 调度规则。
	Pattern []string `json:"pattern"`
}

func (o CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties struct{}"
	}

	return strings.Join([]string{"CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties", string(data)}, " ")
}
