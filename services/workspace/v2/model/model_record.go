package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Record struct {

	// 计算机名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 终端MAC地址。
	TerminalMac *string `json:"terminal_mac,omitempty"`

	// 终端名称。
	TerminalName *string `json:"terminal_name,omitempty"`

	// 终端IP。
	TerminalIp *string `json:"terminal_ip,omitempty"`

	// AccessClient版本。
	ClientVersion *string `json:"client_version,omitempty"`

	// 终端系统类型。
	TerminalType *string `json:"terminal_type,omitempty"`

	// AccessAgent版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 桌面IP。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 开始连接时间。
	ConnectionStartTime *string `json:"connection_start_time,omitempty"`

	// 建立连接时间。
	ConnectionSetupTime *string `json:"connection_setup_time,omitempty"`

	// 结束连接时间。
	ConnectionEndTime *string `json:"connection_end_time,omitempty"`

	// 是否重连。
	IsReconnect *bool `json:"is_reconnect,omitempty"`

	// 连接失败原因。
	ConnectionFailureReason *string `json:"connection_failure_reason,omitempty"`
}

func (o Record) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Record struct{}"
	}

	return strings.Join([]string{"Record", string(data)}, " ")
}
