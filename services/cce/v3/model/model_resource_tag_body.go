package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTagBody 获取到的资源标签
type ResourceTagBody struct {

	// 资源标签key
	Key *string `json:"key,omitempty"`

	// 资源标签value
	Value *string `json:"value,omitempty"`

	// 资源id
	ResourceId *string `json:"resourceId,omitempty"`

	// 资源值列表
	Values *[]string `json:"values,omitempty"`
}

func (o ResourceTagBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagBody struct{}"
	}

	return strings.Join([]string{"ResourceTagBody", string(data)}, " ")
}
