package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpExtendProperties 云外资产IP扩展信息
type OcaIpExtendProperties struct {
	Device *OcaIpDevice `json:"device,omitempty"`

	System *OcaIpSystem `json:"system,omitempty"`

	// 应用信息
	Services *[]OcaIpService `json:"services,omitempty"`
}

func (o OcaIpExtendProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpExtendProperties struct{}"
	}

	return strings.Join([]string{"OcaIpExtendProperties", string(data)}, " ")
}
