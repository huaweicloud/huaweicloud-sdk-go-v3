package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudConnectionDomainBandwidthCapabilityValue struct {
	Bandwidth *CloudConnectionDomainBandwidthValue `json:"bandwidth"`
}

func (o CloudConnectionDomainBandwidthCapabilityValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionDomainBandwidthCapabilityValue struct{}"
	}

	return strings.Join([]string{"CloudConnectionDomainBandwidthCapabilityValue", string(data)}, " ")
}
