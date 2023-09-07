package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEnvironmentRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *EnvironmentKindObj `json:"kind"`

	Metadata *CreateEnvironmentRequestBodyMetadata `json:"metadata"`
}

func (o CreateEnvironmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequestBody", string(data)}, " ")
}
