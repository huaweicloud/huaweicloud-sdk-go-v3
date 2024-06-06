package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetRecyclePolicyRequestBody struct {
	RecyclePolicy *RecyclePolicy `json:"recycle_policy,omitempty"`
}

func (o SetRecyclePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyRequestBody", string(data)}, " ")
}
