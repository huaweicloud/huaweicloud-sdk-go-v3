package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteActionRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ActionKindObj `json:"kind"`

	Metadata *ExecuteActionRequestBodyMetadata `json:"metadata,omitempty"`

	Spec *ActionOnComponentSpec `json:"spec,omitempty"`
}

func (o ExecuteActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteActionRequestBody", string(data)}, " ")
}
