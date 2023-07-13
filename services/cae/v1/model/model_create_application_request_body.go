package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateApplicationRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Application”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *CreateApplicationRequestBodyMetadata `json:"metadata"`
}

func (o CreateApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBody", string(data)}, " ")
}
