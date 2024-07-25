package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResInstanceBody 获取到的资源实例
type ResInstanceBody struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源详情
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// 资源标签
	Tags *[]ResourceTagBody `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]ResourceTagBody `json:"sys_tags,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResInstanceBody struct{}"
	}

	return strings.Join([]string{"ResInstanceBody", string(data)}, " ")
}
