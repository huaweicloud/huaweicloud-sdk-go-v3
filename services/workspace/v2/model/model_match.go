package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Match 根据Tag搜索时的字段结构体。
type Match struct {

	// 搜索时要匹配的字段
	Key string `json:"key"`

	// 搜索时匹配的值，当key=resource_name时为模糊匹配
	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
