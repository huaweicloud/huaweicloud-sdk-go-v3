package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppConnectionInfo 应用连接信息。
type AppConnectionInfo struct {

	// 应用连接唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 应用服务器sid。
	Sid *string `json:"sid,omitempty"`

	// 应用id。
	AppId *string `json:"app_id,omitempty"`

	// 应用服务器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 登录用户。
	UserName *string `json:"user_name,omitempty"`

	// 应用组名称。
	AppGroupName *string `json:"app_group_name,omitempty"`

	// 应用组ID。
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 登录应用时间。
	BrokeringTime *sdktime.SdkTime `json:"brokering_time,omitempty"`

	// 连接失败状态码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 连接失败原因。
	ConnectionFailureReason *string `json:"connection_failure_reason,omitempty"`

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

	// 连接IP。
	WiIp *string `json:"wi_ip,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 会话虚拟ip。
	VirtualIp *string `json:"virtual_ip,omitempty"`

	// 客户端出口ip。
	PublicIp *string `json:"public_ip,omitempty"`

	// 事务id。
	TransactionId *string `json:"transaction_id,omitempty"`

	// 登录应用结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// aps服务器ID。
	ApsInstanceId *string `json:"aps_instance_id,omitempty"`

	// aps服务器名称。
	ApsInstanceName *string `json:"aps_instance_name,omitempty"`

	// wdh专属主机ID。
	ApsHostId *string `json:"aps_host_id,omitempty"`

	// 主服务器组ID。
	PrimaryServerGroupId *string `json:"primary_server_group_id,omitempty"`

	// 主服务器组ID。
	SecondaryServerGroupId *string `json:"secondary_server_group_id,omitempty"`
}

func (o AppConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppConnectionInfo struct{}"
	}

	return strings.Join([]string{"AppConnectionInfo", string(data)}, " ")
}
