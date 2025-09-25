package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckHostInfo 服务器信息
type SecurityCheckHostInfo struct {

	// **参数解释**： 服务器名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器ID **取值范围**： 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 私有IP地址 **取值范围**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 公网IP地址 **取值范围**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 系统类型 **取值范围**： 不涉及
	OsType *string `json:"os_type,omitempty"`
}

func (o SecurityCheckHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckHostInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckHostInfo", string(data)}, " ")
}
