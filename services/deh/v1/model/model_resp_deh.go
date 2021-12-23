package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// resource字段数据结构。
type RespDeh struct {
	// 专属主机ID。

	ResourceId string `json:"resource_id"`
	// 专属主机详情。  该字段用于后续扩展，默认为空。

	ResouceDetail string `json:"resouce_detail"`
	// 标签列表。

	Tags []ResourceTag `json:"tags"`
	// 资源名称。

	ResourceName string `json:"resource_name"`
}

func (o RespDeh) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespDeh struct{}"
	}

	return strings.Join([]string{"RespDeh", string(data)}, " ")
}
