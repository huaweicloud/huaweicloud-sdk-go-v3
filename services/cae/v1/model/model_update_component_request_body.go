package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateComponentRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Component”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *UpdateComponentRequestMetadata `json:"metadata,omitempty"`

	Spec *UpdateComponentRequestSpec `json:"spec,omitempty"`
}

func (o UpdateComponentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestBody", string(data)}, " ")
}
