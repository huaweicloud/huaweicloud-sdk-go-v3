package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppConnectionReq 请求应用使用记录响应体。
type ListAppConnectionReq struct {

	// 主键。
	Id *string `json:"id,omitempty"`

	// 应用服务器sid。
	Sid *string `json:"sid,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 登录用户，模糊查询。
	UserName *string `json:"user_name,omitempty"`

	// 应用组名称。
	AppGroupName *string `json:"app_group_name,omitempty"`

	// 应用组id。
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用名称，模糊查询。
	AppName *string `json:"app_name,omitempty"`

	// 连接失败状态码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 连接失败原因。
	ConnectionFailureReason *string `json:"connection_failure_reason,omitempty"`

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

	// 连接IP。
	WiIp *string `json:"wi_ip,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 登录应用开始时间，格式 2022-10-31 08:07:39。
	BrokeringStartTime *sdktime.SdkTime `json:"brokering_start_time,omitempty"`

	// 登录应用结束时间，格式 2022-10-31 08:07:39。
	BrokeringEndTime *sdktime.SdkTime `json:"brokering_end_time,omitempty"`

	// 会话虚拟ip。
	VirtualIp *string `json:"virtual_ip,omitempty"`
}

func (o ListAppConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppConnectionReq struct{}"
	}

	return strings.Join([]string{"ListAppConnectionReq", string(data)}, " ")
}
