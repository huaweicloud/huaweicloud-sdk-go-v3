package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Match struct {

	// 键。第一期限定为resource_name,后续扩展。
	Key string `json:"key"`

	// 值。每个值最大长度255个字符。
	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
