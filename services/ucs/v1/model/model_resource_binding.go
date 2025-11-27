package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceBinding struct {

	// API类型，固定值“ResourceBinding”
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *ResourceBindingSpec `json:"spec,omitempty"`

	Status *ResourceBindingStatus `json:"status,omitempty"`
}

func (o ResourceBinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBinding struct{}"
	}

	return strings.Join([]string{"ResourceBinding", string(data)}, " ")
}
