package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Event struct {

	// 云主机id
	HostId *string `json:"host_id,omitempty"`

	// 云主机名称
	HostName *string `json:"host_name,omitempty"`

	// 事件类型，包含如下:   - Abnormal Login : 账户异常登录   - Invalid System Account : 风险账号   - Brute Force Cracking : 账号暴力破解   - System Start Script Change : 自启动检测   - Process Abnormal Activity : 进程异常行为   - Process Privilege Escalation : 进程提权操作   - File Privilege Escalation : 文件提权操作   - General Malware : 恶意程序（云查杀）   - Abnormal Shell : 异常shell   - Reverse Shell : 反弹Shell   - High-Risk Command Execution : 高危命令执行   - Key File Change : 关键文件变更   - Webshell : 网站后门
	EventType *string `json:"event_type,omitempty"`

	// 发生时间，毫秒
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 处理时间，毫秒
	HandleTime *int64 `json:"handle_time,omitempty"`

	// 处理状态，包含如下类型：   - \"unhandled\"：未处理   - \"handled\"：已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 处理方式，包含如下类型：   - \"mark_as_handled\" ： 手动处理   - \"ignore\" ： 忽略   - \"add_to_alarm_whitelist\" ：加入告警白名单   - \"add_to_login_whitelist\" ：加入登录白名单   - \"isolate_and_kill\" ：隔离查杀
	HandleMethod *string `json:"handle_method,omitempty"`

	// 事件详细信息，json格式
	AppendInfo *interface{} `json:"append_info,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
