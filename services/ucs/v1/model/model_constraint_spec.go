package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConstraintSpec struct {

	// 策略执行行为，当前支持warn和deny
	EnforcementAction *string `json:"enforcementAction,omitempty"`

	Match *Match `json:"match,omitempty"`

	// 可变参数
	Parameters *interface{} `json:"parameters,omitempty"`
}

func (o ConstraintSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstraintSpec struct{}"
	}

	return strings.Join([]string{"ConstraintSpec", string(data)}, " ")
}
