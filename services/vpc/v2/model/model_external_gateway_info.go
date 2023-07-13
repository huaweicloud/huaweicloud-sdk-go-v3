package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalGatewayInfo
type ExternalGatewayInfo struct {

	// 功能说明：是否启用SNAT 取值范围：true、false；默认为false。
	EnableSnat *bool `json:"enable_snat,omitempty"`

	// 外部网络的ID。
	NetworkId *string `json:"network_id,omitempty"`
}

func (o ExternalGatewayInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalGatewayInfo struct{}"
	}

	return strings.Join([]string{"ExternalGatewayInfo", string(data)}, " ")
}
