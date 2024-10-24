package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioSelectorPidSelection PID选择器
type AudioSelectorPidSelection struct {

	// 设置PID的值
	Pid int32 `json:"pid"`
}

func (o AudioSelectorPidSelection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioSelectorPidSelection struct{}"
	}

	return strings.Join([]string{"AudioSelectorPidSelection", string(data)}, " ")
}
