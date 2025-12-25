package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchPolicyRequestBodyConditionConditions struct {

	// 条件名称
	Name *string `json:"name,omitempty"`

	// 条件键值对
	Data *[]string `json:"data,omitempty"`
}

func (o SearchPolicyRequestBodyConditionConditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBodyConditionConditions struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBodyConditionConditions", string(data)}, " ")
}
