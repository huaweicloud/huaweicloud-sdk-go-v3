package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesClientAutocloseMonitorOptions 自动关闭显示器控制的选项。
type PoliciesClientAutocloseMonitorOptions struct {

	// 自动关闭显示器等待时间。取值范围为[10-600000]。默认：300。
	AutocloseMonitorWaitTime *int32 `json:"autoclose_monitor_wait_time,omitempty"`
}

func (o PoliciesClientAutocloseMonitorOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesClientAutocloseMonitorOptions struct{}"
	}

	return strings.Join([]string{"PoliciesClientAutocloseMonitorOptions", string(data)}, " ")
}
