package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEventsRequest struct {
	// 查询时间段的起始时间，毫秒级时间戳，end_time减去begin_time小于等于2天

	BeginTime string `json:"begin_time"`
	// 查询时间段的终止时间，毫秒级时间戳，end_time减去begin_time小于等于2天

	EndTime string `json:"end_time"`
	// 云主机名称

	HostName *string `json:"host_name,omitempty"`
	// 事件类型，包含如下:   - Abnormal Login : 账户异常登录   - Invalid System Account : 风险账号   - Brute Force Cracking : 账号暴力破解   - System Start Script Change : 自启动检测   - Process Abnormal Activity : 进程异常行为   - Process Privilege Escalation : 进程提权操作   - File Privilege Escalation : 文件提权操作   - General Malware : 恶意程序（云查杀）   - Abnormal Shell : 异常shell   - Reverse Shell : 反弹Shell   - High-Risk Command Execution : 高危命令执行   - Key File Change : 关键文件变更   - Webshell : 网站后门

	EventTypes []string `json:"event_types"`
	// 是否已处理，包含如下类型：   - \"unhandled\" ： 未处理   - \"handled\" ： 已处理

	HandleStatus *string `json:"handle_status,omitempty"`
	// 默认10

	Limit *int32 `json:"limit,omitempty"`
	// 默认0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}
