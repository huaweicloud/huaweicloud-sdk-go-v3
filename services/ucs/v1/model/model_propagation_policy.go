package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropagationPolicy struct {

	// API类型，固定值“PropagationPolicy”
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *PropagationSpec `json:"spec,omitempty"`
}

func (o PropagationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropagationPolicy struct{}"
	}

	return strings.Join([]string{"PropagationPolicy", string(data)}, " ")
}
