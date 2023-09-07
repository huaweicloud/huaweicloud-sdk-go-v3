package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ComponentKindObj `json:"kind"`

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
