package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateNetworkSegmentsVo **参数解释**： 私网网段信息列表 **取值范围**： 不涉及
type ListPrivateNetworkSegmentsVo struct {

	// **参数解释**： 私网网段列表 **取值范围**： 不涉及
	PrivateNetworkList *[]PrivateNetworkSegmentVo `json:"private_network_list,omitempty"`

	// **参数解释**： 总条数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 私网网段配额 **取值范围**： 不涉及
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**： 支持的可用区列表 **取值范围**： 不涉及
	SupportAzList *[]string `json:"support_az_list,omitempty"`
}

func (o ListPrivateNetworkSegmentsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNetworkSegmentsVo struct{}"
	}

	return strings.Join([]string{"ListPrivateNetworkSegmentsVo", string(data)}, " ")
}
