package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionInfo 逻辑集群操作信息
type ActionInfo struct {

	// 操作名称。当前只允许Create,Expand,Restart,Delete,Shrink
	ActionName *string `json:"action_name,omitempty"`

	// 操作进度，默认10
	Progress *int32 `json:"progress,omitempty"`

	// 是否完成操作
	Completed *bool `json:"completed,omitempty"`

	// 操作开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 操作结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 操作结果。success或者failed，默认空字符串
	Result *string `json:"result,omitempty"`

	// 操作日志
	Logs *string `json:"logs,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}
