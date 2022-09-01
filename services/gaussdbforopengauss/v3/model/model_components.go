package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件列表
type Components struct {

	// 组件id，当组件类型为DN，组件id为6001，则对应的值为dn_6001。
	Id *string `json:"id,omitempty" xml:"id"`

	// 节点类型，取值为“master”、“slave”，分别对应于主节点、备节点。
	Role *string `json:"role,omitempty" xml:"role"`

	// 节点状态。
	Status *string `json:"status,omitempty" xml:"status"`
}

func (o Components) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Components struct{}"
	}

	return strings.Join([]string{"Components", string(data)}, " ")
}
