package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDecoyPortPolicyDetailsResponse Response Object
type ShowDecoyPortPolicyDetailsResponse struct {

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 端口与协议
	PortList *[]PolicyDetailsPortList `json:"port_list,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// ip白名单
	WhiteIp *[]string `json:"white_ip,omitempty"`

	// 主机列表
	HostList       *[]string `json:"host_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDecoyPortPolicyDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDecoyPortPolicyDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDecoyPortPolicyDetailsResponse", string(data)}, " ")
}
