package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepCondition 条件节点执行条件。
type StepCondition struct {

	// 判断类型，例如==（等于）、!=（不等于）、>（大于）、>=（大于等于）、<（小于）、<=（小于等于）、in（包含）、or（或）。
	Type *string `json:"type,omitempty"`

	// 节点执行条件为true时的分支。
	Left *interface{} `json:"left,omitempty"`

	// 节点执行条件为false时的分支。
	Right *interface{} `json:"right,omitempty"`
}

func (o StepCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepCondition struct{}"
	}

	return strings.Join([]string{"StepCondition", string(data)}, " ")
}
