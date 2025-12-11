package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusResultResponseInfo 病毒查杀结果列表响应详情
type AntiVirusResultResponseInfo struct {

	// **参数解释**： 病毒查杀结果ID **取值范围**： 字符长度1-64位
	ResultId *string `json:"result_id,omitempty"`

	// **参数解释**： 病毒类型 **取值范围**： Trojan（木马）、Virus（病毒）、Worm（蠕虫）等
	MalwareType *string `json:"malware_type,omitempty"`

	// **参数解释**： 病毒名称 **取值范围**： 字符长度1-128位
	MalwareName *string `json:"malware_name,omitempty"`

	// **参数解释**: 威胁等级 **取值范围**: 包含如下:   - Low：低危   - Medium：中危   - High：高危   - Critical：致命
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 任务ID **取值范围**: 字符长度1-64位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 任务名称 **取值范围**: 最大长度255个unicode字符。
	TaskName *string `json:"task_name,omitempty"`

	FileInfo *ResultFileResponseInfo `json:"file_info,omitempty"`

	ResourceInfo *ResultResourceResponseInfo `json:"resource_info,omitempty"`

	// **参数解释**: 病毒查杀结果对应的事件类型标识 **取值范围**: 0-10（具体含义：0=文件病毒事件、1=内存病毒事件...，详见产品错误码/枚举文档）
	EventType *int32 `json:"event_type,omitempty"`

	// **参数解释**： 发生时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 处理状态 **取值范围**： - unhandled：未处理 - handled：已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 处理方式 **取值范围**: 包含如下:   - mark_as_handled：手动处理   - ignore：忽略   - add_to_alarm_whitelist：加入告警白名单   - isolate_and_kill：隔离文件   - unhandle：取消手动处理   - do_not_ignore：取消忽略   - remove_from_alarm_whitelist：删除告警白名单   - do_not_isolate_or_kill：取消隔离文件
	HandleMethod *string `json:"handle_method,omitempty"`

	// **参数解释** 备注信息 **取值范围** 字符长度0-512位
	Memo *string `json:"memo,omitempty"`

	// **参数解释**: 后续处置操作列表 **取值范围**: 数组元素为处置操作枚举字符串（如“isolate_and_kill”“ignore”等），数组长度0-4（具体支持操作因结果状态而异）
	OperateAcceptList *[]string `json:"operate_accept_list,omitempty"`

	// **参数解释**: 操作详情信息列表（页面不展示） **取值范围**: 数组长度0-100
	OperateDetailList *[]ResultDetailResponseInfo `json:"operate_detail_list,omitempty"`

	// **参数解释**: 自动隔离查杀标识 **取值范围**: 字符长度1-16位，枚举值为“auto_isolate”（自动隔离）、“manual”（手动操作）、“none”（未隔离）
	IsolateTag *string `json:"isolate_tag,omitempty"`
}

func (o AntiVirusResultResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusResultResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusResultResponseInfo", string(data)}, " ")
}
