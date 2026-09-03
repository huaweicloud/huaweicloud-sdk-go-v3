package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyReqBody struct {

	// 策略的更新描述。
	Description *string `json:"description,omitempty"`

	Definition *PolicyDefinition `json:"definition,omitempty"`

	ValidationMode *ValidationMode `json:"validation_mode,omitempty"`
}

func (o UpdatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyReqBody", string(data)}, " ")
}
