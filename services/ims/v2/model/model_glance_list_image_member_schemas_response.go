package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceListImageMemberSchemasResponse Response Object
type GlanceListImageMemberSchemasResponse struct {

	// 视图链接。
	Links *[]Links `json:"links,omitempty"`

	// 视图名称。
	Name *string `json:"name,omitempty"`

	// 镜像属性说明，主要是对基础属性的说明，包含每个属性的取值类型，用途。
	Properties     *interface{} `json:"properties,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GlanceListImageMemberSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMemberSchemasResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageMemberSchemasResponse", string(data)}, " ")
}
