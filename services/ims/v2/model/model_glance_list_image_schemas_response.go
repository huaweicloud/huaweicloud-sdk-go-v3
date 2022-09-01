package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceListImageSchemasResponse struct {

	// 视图名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 镜像属性说明，主要是对基础属性的说明，包含每个属性的取值类型、用途等。
	Properties *interface{} `json:"properties,omitempty" xml:"properties"`

	// 视图链接。
	Links          *[]Links `json:"links,omitempty" xml:"links"`
	HttpStatusCode int      `json:"-"`
}

func (o GlanceListImageSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasResponse", string(data)}, " ")
}
