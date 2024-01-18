package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserConnectionInfo 用户登录记录信息。
type UserConnectionInfo struct {

	// 主键。
	Id *string `json:"id,omitempty"`

	// 连接类型。
	ConnectType *string `json:"connect_type,omitempty"`

	// 登录用户。
	UserName *string `json:"user_name,omitempty"`

	// 桌面组名。
	DesktopGroupName *string `json:"desktop_group_name,omitempty"`

	// 预连接时间。
	PreConnTime *sdktime.SdkTime `json:"pre_conn_time,omitempty"`

	// 开始时间。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 应用服务器sid。
	MachineSid *string `json:"machine_sid,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 连接失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`

	// 连接失败状态码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 客户端Mac。
	ClientMac *string `json:"client_mac,omitempty"`

	// 客户端名称。
	ClientName *string `json:"client_name,omitempty"`

	// 客户端ip。
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端版本。
	ClientVersion *string `json:"client_version,omitempty"`

	// 客户端操作系统类型。
	ClientType *string `json:"client_type,omitempty"`

	// aps hda版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 应用服务器ip。
	VmIp *string `json:"vm_ip,omitempty"`

	// 连接标志。
	ConnectFlag *string `json:"connect_flag,omitempty"`

	// 连接IP。
	WiIp *string `json:"wi_ip,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 会话虚拟ip。
	VirtualIp *string `json:"virtual_ip,omitempty"`
}

func (o UserConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserConnectionInfo struct{}"
	}

	return strings.Join([]string{"UserConnectionInfo", string(data)}, " ")
}
