package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopAction 桌面开关机实体
type DesktopAction struct {

	// 行为动作
	Action *string `json:"action,omitempty"`

	// 行为完成状态信息
	Message *string `json:"message,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 结果
	Result *string `json:"result,omitempty"`

	// 异常信息
	Traceback *string `json:"traceback,omitempty"`
}

func (o DesktopAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopAction struct{}"
	}

	return strings.Join([]string{"DesktopAction", string(data)}, " ")
}
