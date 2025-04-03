package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionConfig 动作编排配置。
type ActionConfig struct {

	// 算法自动插入无语义动作的时间间隔。这个参数填0或者不填默认是间隔4秒，设置成255时不自动插入无语义动作。
	ActionInterval *float32 `json:"action_interval,omitempty"`
}

func (o ActionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionConfig struct{}"
	}

	return strings.Join([]string{"ActionConfig", string(data)}, " ")
}
