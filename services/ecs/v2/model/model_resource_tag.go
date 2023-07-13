package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 云服务器标签。
type ResourceTag struct {

	// 键。  - 最大长度127个unicode字符。  - key不能为空。  - 只能包含字母、数字、下划线“_”、中划线“-”。
	Key string `json:"key"`

	// 值。  - 每个值最大长度255个unicode字符。  - 可以为空字符串。  - 只能包含字母、数字、下划线“_”、中划线“-”。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
