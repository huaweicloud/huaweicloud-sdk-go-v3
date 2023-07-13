package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteActionRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Action”，该值不可修改。
	Kind string `json:"kind"`

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
