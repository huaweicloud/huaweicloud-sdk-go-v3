package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo **参数解释**： 防护配置对应的主机信息 **取值范围**： 不涉及
type WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo struct {

	// **参数解释**: 主机id **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机私网ip **取值范围**: 字符长度1-128位
	HostPrivateIp *string `json:"host_private_ip,omitempty"`
}

func (o WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionConfigResponseInfoContainerWtpInfoHostInfo", string(data)}, " ")
}
