package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Rule struct {

	// API类型。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *RuleSpec `json:"spec,omitempty"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
