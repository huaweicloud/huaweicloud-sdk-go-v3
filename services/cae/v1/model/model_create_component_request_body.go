package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Component”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *CreateComponentRequestBodyMetadata `json:"metadata,omitempty"`

	Spec *CreateComponentRequestBodySpec `json:"spec,omitempty"`
}

func (o CreateComponentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComponentRequestBody", string(data)}, " ")
}
