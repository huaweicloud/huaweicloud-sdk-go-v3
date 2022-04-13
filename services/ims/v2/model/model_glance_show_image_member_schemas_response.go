package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceShowImageMemberSchemasResponse struct {
	// 视图名称。

	Name *string `json:"name,omitempty"`
	// 镜像成员属性说明，主要是对基础属性的说明，包含每个属性的取值类型、用途等。

	Properties     *interface{} `json:"properties,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GlanceShowImageMemberSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberSchemasResponse", string(data)}, " ")
}
