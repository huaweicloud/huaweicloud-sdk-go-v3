package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusTaskRequestInfo 创建扫描任务
type CreateAntiVirusTaskRequestInfo struct {

	// 任务名称
	TaskName string `json:"task_name"`

	// 任务ID 创建病毒扫描任务时,task_id是null.重新扫描时，task_id不是null,是当前任务的ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType string `json:"scan_type"`

	// 处置动作，包含如下:   - auto：自动处置   - manual：人工处置
	Action string `json:"action"`

	// 策略管理主机列表
	HostIds []string `json:"host_ids"`
}

func (o CreateAntiVirusTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusTaskRequestInfo", string(data)}, " ")
}
