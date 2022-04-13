package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyRequestBody struct {
	// 防护策略名

	Name *string `json:"name,omitempty"`

	Action *PolicyAction `json:"action,omitempty"`

	Options *PolicyOption `json:"options,omitempty"`
}

func (o UpdatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequestBody", string(data)}, " ")
}
