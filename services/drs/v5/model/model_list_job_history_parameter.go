package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobHistoryParameter 任务参数配置历史修改响应体
type ListJobHistoryParameter struct {

	// 参数名称。
	Name string `json:"name"`

	// 旧参数值。
	OldValue string `json:"old_value"`

	// 新参数值。
	NewValue string `json:"new_value"`

	// 更新结果。true：成功，false：失败
	IsUpdateSuccess bool `json:"is_update_success"`

	// 是否已应用。true：已应用，false：未应用
	IsApplied bool `json:"is_applied"`

	// 参数修改时间。
	UpdateTime string `json:"update_time"`

	// 参数应用时间。
	ApplyTime *string `json:"apply_time,omitempty"`
}

func (o ListJobHistoryParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobHistoryParameter struct{}"
	}

	return strings.Join([]string{"ListJobHistoryParameter", string(data)}, " ")
}
