package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 根据标签查询资源时，match字段结构。
type ActionMatch struct {
	// 键。目前只支持resource_name。表示匹配资源实例的名称。

	Key string `json:"key"`
	// 值。每个值最大长度255个unicode字符 。当key为resource_name时，表示待匹配的资源实例的名称。

	Value string `json:"value"`
}

func (o ActionMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionMatch struct{}"
	}

	return strings.Join([]string{"ActionMatch", string(data)}, " ")
}
