package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tags
type Tags struct {

	// 键。最大长度128个unicode字符。搜索时不对此参数做校验，key不能为空或者空字符串，不能为空格，校验和使用之前先trim 前后空格。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个unicode字符。
	Values []string `json:"values"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
