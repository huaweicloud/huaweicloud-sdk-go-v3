package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务参数校验参数
type Constraint struct {
	// 校验规则类型

	Type string `json:"type"`
	// 校验规则

	Value string `json:"value"`
	// 校验失败描述

	Errormsg string `json:"errormsg"`
}

func (o Constraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Constraint struct{}"
	}

	return strings.Join([]string{"Constraint", string(data)}, " ")
}
