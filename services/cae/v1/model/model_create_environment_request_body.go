package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEnvironmentRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Environment”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *CreateEnvironmentRequestBodyMetadata `json:"metadata"`
}

func (o CreateEnvironmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequestBody", string(data)}, " ")
}
