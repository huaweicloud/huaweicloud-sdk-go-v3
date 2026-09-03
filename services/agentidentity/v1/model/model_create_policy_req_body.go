package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePolicyReqBody struct {

	// Human-readable display name for the policy
	Name string `json:"name"`

	// 策略的可读描述。
	Description *string `json:"description,omitempty"`

	Definition *PolicyDefinition `json:"definition"`

	ValidationMode *ValidationMode `json:"validation_mode,omitempty"`
}

func (o CreatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyReqBody", string(data)}, " ")
}
