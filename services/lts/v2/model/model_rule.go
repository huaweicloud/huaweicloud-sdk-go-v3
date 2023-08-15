package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Rule struct {

	// 内置系统模板类型。
	Type string `json:"type"`

	// 系统iton模板名称。
	Param string `json:"param"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
