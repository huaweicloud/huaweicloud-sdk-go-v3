package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Match 匹配字段
type Match struct {

	// 键。有matches参数时，该字段为必填，固定为resource_name。
	Key string `json:"key"`

	// 值。即资源名称，有matches参数时，该字段为必填，且默认为模糊搜索，如”message.com”。每个值最大长度255个字符。
	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
