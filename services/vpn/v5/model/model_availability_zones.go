package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZones 可用区情况
type AvailabilityZones struct {
	Basic *VpnGatewayAvailabilityZones `json:"basic,omitempty"`

	Professional1 *VpnGatewayAvailabilityZones `json:"professional1,omitempty"`

	Professional2 *VpnGatewayAvailabilityZones `json:"professional2,omitempty"`

	Gm *VpnGatewayAvailabilityZones `json:"gm,omitempty"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}
