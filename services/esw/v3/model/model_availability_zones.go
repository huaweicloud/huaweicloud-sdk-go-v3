package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailabilityZones struct {

	// - 参数解释：ESW实例默认主节点所在的可用区。 - 约束限制：1-128字符。 - 取值范围：当前区域可用区id。 - 默认取值：不涉及。
	Primary string `json:"primary"`

	// - 参数解释：ESW实例默认备节点所在的可用区。 - 约束限制：1-128字符。 - 取值范围：当前区域可用区id。 - 默认取值：不涉及。
	Standby string `json:"standby"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}
