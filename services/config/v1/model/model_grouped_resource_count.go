package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupedResourceCount struct {

	// 分组名称。
	GroupName *string `json:"group_name,omitempty"`

	// 资源数量。
	ResourceCount *int32 `json:"resource_count,omitempty"`
}

func (o GroupedResourceCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupedResourceCount struct{}"
	}

	return strings.Join([]string{"GroupedResourceCount", string(data)}, " ")
}
