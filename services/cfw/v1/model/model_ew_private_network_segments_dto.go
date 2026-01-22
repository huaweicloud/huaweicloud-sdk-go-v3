package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EwPrivateNetworkSegmentsDto struct {

	// **参数解释**： 私网网段的信息列表，用于东西向VPC防护引流 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PrivateNetworkSegments []PrivateNetworkSegmentVo `json:"private_network_segments"`
}

func (o EwPrivateNetworkSegmentsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EwPrivateNetworkSegmentsDto struct{}"
	}

	return strings.Join([]string{"EwPrivateNetworkSegmentsDto", string(data)}, " ")
}
