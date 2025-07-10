package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddServerTag 云服务器标签。
type BatchAddServerTag struct {

	// 键。  - 不能为空。  - 对于同一资源键值唯一。  - 长度不超过36个字符。  - 字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Key string `json:"key"`

	// 值。  - 长度不超过43个字符。  - 字符集：A-Z，a-z ， 0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。  - 只能包含数字、字母、中划线“-”、下划线“_”。
	Value *string `json:"value,omitempty"`
}

func (o BatchAddServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerTag struct{}"
	}

	return strings.Join([]string{"BatchAddServerTag", string(data)}, " ")
}
