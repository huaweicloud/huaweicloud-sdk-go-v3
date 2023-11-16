package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Rule struct {

	// 规则ID
	Id string `json:"id"`

	// 规则类型
	Type string `json:"type"`

	// 规则名称
	Name string `json:"name"`

	// 规则版本
	Version string `json:"version"`

	// 最近操作人员
	Operator string `json:"operator"`

	// 最近操作时间
	OperateTime int64 `json:"operate_time"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
