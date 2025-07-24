package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypeUsage struct {
	Capacity *ShareTypeUsageCapacity `json:"capacity,omitempty"`

	Bandwidth *ShareTypeUsageBandwidth `json:"bandwidth,omitempty"`

	Quantity *ShareTypeUsageQuantity `json:"quantity,omitempty"`
}

func (o ShareTypeUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeUsage struct{}"
	}

	return strings.Join([]string{"ShareTypeUsage", string(data)}, " ")
}
