package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleId struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`
}

func (o RuleId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleId struct{}"
	}

	return strings.Join([]string{"RuleId", string(data)}, " ")
}
