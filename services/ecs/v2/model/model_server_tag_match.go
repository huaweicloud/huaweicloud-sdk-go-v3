package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerTagMatch 搜索字段，用于按条件搜索弹性云服务器。
type ServerTagMatch struct {

	// 键，表示要匹配的字段。  当前key的参数值只能取“resource_name”，此时value的参数值为云服务器名称。  - key不能重复，value为匹配的值。  - 此字段为固定字典值。  - 不允许为空字符串。
	Key string `json:"key"`

	// 值。  当前key的参数值只能取“resource_name”，此时value的参数值为云服务器名称。  - 每个值最大长度255个unicode字符。  - 不可以为空。
	Value string `json:"value"`
}

func (o ServerTagMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerTagMatch struct{}"
	}

	return strings.Join([]string{"ServerTagMatch", string(data)}, " ")
}
