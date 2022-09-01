package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源类型：server(服务器)
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
