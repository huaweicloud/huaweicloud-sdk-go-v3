package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SessionInfo 会话信息。
type SessionInfo struct {

	// 主键ID。
	Id *string `json:"id,omitempty"`

	// 会话标识。
	SessionStamp *string `json:"session_stamp,omitempty"`

	// 会话在hda的os中会话id。
	OsSessionId *string `json:"os_session_id,omitempty"`

	// 协议类型。
	ProtocolType *string `json:"protocol_type,omitempty"`

	// 当前会话的登录用户。
	LoginUser *string `json:"login_user,omitempty"`

	// 会话类型，1表示共享桌面，2表示应用。
	SessionType *string `json:"session_type,omitempty"`

	// AppServer组ID。
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// AppServer组名称。
	AppServerGroupName *string `json:"app_server_group_name,omitempty"`

	// 预连接时间。
	PreConnTime *sdktime.SdkTime `json:"pre_conn_time,omitempty"`

	// 会话开始时间。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 会话结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 状态持续时间。
	StatusContinueTime *string `json:"status_continue_time,omitempty"`

	// 服务器SID。
	MachineSid *string `json:"machine_sid,omitempty"`

	// 服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 会话状态。
	SessionState *string `json:"session_state,omitempty"`

	// 会话中的应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 客户端Mac地址。
	ClientMac *string `json:"client_mac,omitempty"`

	// 客户端名称。
	ClientName *string `json:"client_name,omitempty"`

	// 客户端IP。
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端版本。
	ClientVersion *string `json:"client_version,omitempty"`

	// 客户端类型。
	ClientType *string `json:"client_type,omitempty"`

	// agent版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 服务器IP。
	VmIp *string `json:"vm_ip,omitempty"`

	// 错误原因消息。
	FailedReason *string `json:"failed_reason,omitempty"`

	// 错误原因码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 状态最后变化时间。
	LastUpdateStatusTime *sdktime.SdkTime `json:"last_update_status_time,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o SessionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionInfo struct{}"
	}

	return strings.Join([]string{"SessionInfo", string(data)}, " ")
}
