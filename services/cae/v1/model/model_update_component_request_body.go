package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateComponentRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ComponentKindObj `json:"kind"`

	Metadata *UpdateComponentRequestMetadata `json:"metadata"`

	Spec *UpdateComponentRequestSpec `json:"spec"`
}

func (o UpdateComponentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestBody", string(data)}, " ")
}
