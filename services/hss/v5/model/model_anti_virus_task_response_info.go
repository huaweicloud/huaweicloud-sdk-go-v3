package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusTaskResponseInfo 扫描任务信息
type AntiVirusTaskResponseInfo struct {

	// **参数解释**： 任务ID **取值范围**: 字符长度1-64位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 任务名称 **取值范围**: 最大长度255个unicode字符。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 任务类型 **取值范围**： 包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType *string `json:"scan_type,omitempty"`

	// **参数解释**： 启动类型 **取值范围**： 包含如下   - now：立即启动   - later：稍后启动   - period：周期启动
	StartType *string `json:"start_type,omitempty"`

	// **参数解释**: 处置动作 **取值范围**: - auto：自动处置 - manual：人工处置
	Action *string `json:"action,omitempty"`

	// **参数解释**： 启动时间 **取值范围**： 最小值0，最大值9223372036854775807；时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算）；单位：ms
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 任务状态 **取值范围**: 包含如下2种   - scanning：扫描中   - finish：扫描完成
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**: 关联服务器数 **取值范围**: 非负整数，最小值0；单位：台
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 扫描成功服务器数 **取值范围**: 非负整数，最小值0；单位：台
	SuccessHostNum *int32 `json:"success_host_num,omitempty"`

	// **参数解释**: 扫描失败服务器数 **取值范围**: 非负整数，最小值0；单位：台
	FailHostNum *int32 `json:"fail_host_num,omitempty"`

	// **参数解释**: 已取消服务器数 **取值范围**: 非负整数，最小值0；单位：台
	CancelHostNum *int32 `json:"cancel_host_num,omitempty"`

	// **参数解释**: 关联的服务器详细信息列表 **取值范围**: 数组个数0-10000
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
