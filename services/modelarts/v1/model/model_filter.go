package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Filter 查询作业要过滤的一系列条件。
type Filter struct {

	// 分组条件键值。
	Key *string `json:"key,omitempty"`

	// 分组条件键值键关系，支持between（范围）、like（类似）、in（包含）、not（非）。
	Operator *string `json:"operator,omitempty"`

	// 分组条件键对应值。
	Value *[]string `json:"value,omitempty"`
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}
