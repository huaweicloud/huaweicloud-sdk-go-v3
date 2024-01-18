package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserConnectionReq 请求用户登录记录响应体。
type ListUserConnectionReq struct {

	// 登录记录唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 用户连接类别。
	ConnectType *string `json:"connect_type,omitempty"`

	// 登录用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 应用服务器sid。
	MachineSid *string `json:"machine_sid,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 连接失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`

	// 连接失败状态码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 客户端名称。
	ClientName *string `json:"client_name,omitempty"`

	// 客户端版本。
	ClientVersion *string `json:"client_version,omitempty"`

	// 客户端操作系统类型。
	ClientType *string `json:"client_type,omitempty"`

	// aps hda版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 应用服务器ip。
	VmIp *string `json:"vm_ip,omitempty"`

	// 连接标志，目前值为0。
	ConnectFlag *string `json:"connect_flag,omitempty"`

	// 连接IP。
	WiIp *string `json:"wi_ip,omitempty"`

	// 登录应用开始时间，格式 2022-10-31 08:07:39。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 登录应用开始时间，格式 2022-10-31 08:07:39。
	LoginStartTime *sdktime.SdkTime `json:"login_start_time,omitempty"`

	// 登录应用结束时间，格式 2022-10-31 08:07:39。
	LoginEndTime *sdktime.SdkTime `json:"login_end_time,omitempty"`

	// 会话虚拟ip。
	VirtualIp *string `json:"virtual_ip,omitempty"`
}

func (o ListUserConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserConnectionReq struct{}"
	}

	return strings.Join([]string{"ListUserConnectionReq", string(data)}, " ")
}
