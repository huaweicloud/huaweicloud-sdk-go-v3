package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePolicyEngineReqBody struct {

	// Customer-assigned immutable name for the policy engine.
	Name string `json:"name"`

	Type *PolicyEngineType `json:"type"`

	// 策略集的可读描述。
	Description *string `json:"description,omitempty"`

	// 自定义标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePolicyEngineReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyEngineReqBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyEngineReqBody", string(data)}, " ")
}
