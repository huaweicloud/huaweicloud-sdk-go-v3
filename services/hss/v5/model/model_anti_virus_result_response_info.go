package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusResultResponseInfo 病毒查杀结果列表响应详情
type AntiVirusResultResponseInfo struct {

	// 病毒查杀结果ID
	ResultId *string `json:"result_id,omitempty"`

	// 病毒类型
	MalwareType *string `json:"malware_type,omitempty"`

	// 病毒名称
	MalwareName *string `json:"malware_name,omitempty"`

	// 威胁等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 致命
	Severity *string `json:"severity,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	FileInfo *ResultFileResponseInfo `json:"file_info,omitempty"`

	ResourceInfo *ResultResourceResponseInfo `json:"resource_info,omitempty"`

	// 事件类型
	EventType *int32 `json:"event_type,omitempty"`

	// 发生时间，毫秒
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 处置状态，包含如下:   - unhandled：未处理   - handled: 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 处理方式，包含如下:   - mark_as_handled : 手动处理   - ignore : 忽略   - add_to_alarm_whitelist : 加入告警白名单   - isolate_and_kill : 隔离文件
	HandleMethod *string `json:"handle_method,omitempty"`

	// 备注信息
	Memo *string `json:"memo,omitempty"`

	// 支持的处理操作
	OperateAcceptList *[]string `json:"operate_accept_list,omitempty"`

	// 操作详情信息列表（页面不展示）
	OperateDetailList *[]ResultDetailResponseInfo `json:"operate_detail_list,omitempty"`

	// 自动隔离查杀标识
	IsolateTag *string `json:"isolate_tag,omitempty"`
}

func (o AntiVirusResultResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusResultResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusResultResponseInfo", string(data)}, " ")
}
