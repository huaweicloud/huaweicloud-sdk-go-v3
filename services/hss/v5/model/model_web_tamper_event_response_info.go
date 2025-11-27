package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperEventResponseInfo **参数解释**: 网页防篡改事件信息 **取值范围**: 不涉及
type WebTamperEventResponseInfo struct {

	// **参数解释**: 防护文件 **取值范围**: 字符长度0-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 事件描述 **取值范围**: 字符长度0-512位
	EventDescription *string `json:"event_description,omitempty"`

	// **参数解释**: 进程路径 **取值范围**: 字符长度0-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释**: 进程命令行 **取值范围**: 字符长度0-256位
	ProcessCmd *string `json:"process_cmd,omitempty"`

	// **参数解释**: 进程pid **取值范围**: 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**: 进程环境，是指在容器内的进程或者宿主机的进程 **取值范围**: 字符长度0-32位
	ProcessEnv *string `json:"process_env,omitempty"`

	// **参数解释**: 网站应用名称 **取值范围**: 字符长度0-128位
	WebAppName *string `json:"web_app_name,omitempty"`

	// **参数解释**: 检测时间(ms) **取值范围**: 最小值0，最大值9223372036854775807
	EventTime *int64 `json:"event_time,omitempty"`

	HostInfo *WebTamperEventHostInfo `json:"host_info,omitempty"`

	ContainerInfo *WebTamperEventContainerInfo `json:"container_info,omitempty"`
}

func (o WebTamperEventResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperEventResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperEventResponseInfo", string(data)}, " ")
}
