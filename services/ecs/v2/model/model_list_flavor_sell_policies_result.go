package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFlavorSellPoliciesResult struct {

	// 云服务器规格的索引。
	Id int32 `json:"id"`

	// 云服务器规格的ID。
	FlavorId string `json:"flavor_id"`

	// 云服务器规格的售卖状态。  sellout：售罄。 available：可用。
	SellStatus string `json:"sell_status"`

	// 云服务器规格的可用区。
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 云服务器规格的付费模式。  - postPaid：按需计费实例。 - prePaid：包年/包月计费实例。 - spot：竞价实例。 - ri：预留实例。
	SellMode string `json:"sell_mode"`

	SpotOptions *FlavorSpotOptions `json:"spot_options,omitempty"`
}

func (o ListFlavorSellPoliciesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorSellPoliciesResult struct{}"
	}

	return strings.Join([]string{"ListFlavorSellPoliciesResult", string(data)}, " ")
}
