package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateApplicationRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ApplicationKindObj `json:"kind"`

	Metadata *CreateApplicationRequestBodyMetadata `json:"metadata"`
}

func (o CreateApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBody", string(data)}, " ")
}
