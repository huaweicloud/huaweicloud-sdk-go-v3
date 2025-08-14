package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusTaskResponseInfo 扫描任务信息
type AntiVirusTaskResponseInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型，包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType *string `json:"scan_type,omitempty"`

	// 启动类型，包含如下:   - now : 立即启动   - later : 稍后启动   - period : 周期启动
	StartType *string `json:"start_type,omitempty"`

	// 处置动作，包含如下:   - auto：自动处置   - manual：人工处置
	Action *string `json:"action,omitempty"`

	// 启动时间，毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务状态，包含如下2种   - scanning ：扫描中   - finish ：扫描完成
	TaskStatus *string `json:"task_status,omitempty"`

	// 关联服务器数
	HostNum *int32 `json:"host_num,omitempty"`

	// 扫描成功服务器数
	SuccessHostNum *int32 `json:"success_host_num,omitempty"`

	// 扫描失败服务器数
	FailHostNum *int32 `json:"fail_host_num,omitempty"`

	// 已取消服务器数
	CancelHostNum *int32 `json:"cancel_host_num,omitempty"`

	// 主机信息
	HostInfoList *[]AntiVirusTaskHostResponseInfo `json:"host_info_list,omitempty"`

	// 是否需要重新扫描
	Rescan *bool `json:"rescan,omitempty"`

	// 此次扫描任务是否付费
	WhetherPaidTask *bool `json:"whether_paid_task,omitempty"`
}

func (o AntiVirusTaskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusTaskResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusTaskResponseInfo", string(data)}, " ")
}
