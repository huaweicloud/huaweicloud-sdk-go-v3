package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceShowImageSchemasResponse struct {
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
	// 视图名称。

	Name *string `json:"name,omitempty"`
	// 镜像属性说明，主要是对基础属性的说明，包含每个属性的取值类型、用途等。

	Properties *interface{} `json:"properties,omitempty"`
	// 视图链接。

	Links          *[]Links `json:"links,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o GlanceShowImageSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageSchemasResponse", string(data)}, " ")
}
