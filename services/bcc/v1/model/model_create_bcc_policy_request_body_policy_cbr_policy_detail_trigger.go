package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger 调度器属性。
type CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger struct {
	Properties *CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTriggerProperties `json:"properties"`
}

func (o CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger struct{}"
	}

	return strings.Join([]string{"CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger", string(data)}, " ")
}
