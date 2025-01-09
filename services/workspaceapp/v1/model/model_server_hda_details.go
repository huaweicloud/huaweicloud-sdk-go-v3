package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerHdaDetails 服务器的accessAgent详细信息。
type ServerHdaDetails struct {

	// 服务器id。
	ServerId *string `json:"server_id,omitempty"`

	// 机器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 是否是维护状态。
	MaintainStatus *bool `json:"maintain_status,omitempty"`

	// 服务器名称。
	ServerName *string `json:"server_name,omitempty"`

	// 服务器组id。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 服务器组名称。
	ServerGroupName *string `json:"server_group_name,omitempty"`

	// 服务器的sid。
	Sid *string `json:"sid,omitempty"`

	// 会话数量。
	SessionCount *int32 `json:"session_count,omitempty"`

	Status *ServerStatus `json:"status,omitempty"`

	// 当前的accessAgent版本。
	CurrentVersion *string `json:"current_version,omitempty"`
}

func (o ServerHdaDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHdaDetails struct{}"
	}

	return strings.Join([]string{"ServerHdaDetails", string(data)}, " ")
}
