package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OverridePolicy struct {

	// API类型，固定值“OverridePolicy”
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Spec *OverrideSpec `json:"spec,omitempty"`
}

func (o OverridePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverridePolicy struct{}"
	}

	return strings.Join([]string{"OverridePolicy", string(data)}, " ")
}
