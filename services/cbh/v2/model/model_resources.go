package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resources 查询资源实例列表返回体中用于存储标签列表的结构体。
type Resources struct {

	// 实例ID。
	ResourceId string `json:"resource_id"`

	ResourceDetail *InstanceDetail `json:"resource_detail"`

	// tags。
	Tags []ResourceTag `json:"tags"`

	// sys_tags。
	SysTags []ResourceTag `json:"sys_tags"`

	// 资源名称。
	ResourceName string `json:"resource_name"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
