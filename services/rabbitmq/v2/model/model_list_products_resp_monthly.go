package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespMonthly struct {

	// **参数解释**： 消息引擎的名称，该字段显示为kafka。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消息引擎的版本。 **取值范围**： [- 3.8.35](tag:hws,hws_hk,hws_eu,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,ax,srg) [- 3.12.13](tag:srg) [- 3.13.7](tag:dt) [- AMQP-0-9-1](tag:hws,hws_hk,hws_eu)
	Version *string `json:"version,omitempty"`

	// **参数解释**： 产品规格列表。
	Values *[]ListProductsRespValues1 `json:"values,omitempty"`
}

func (o ListProductsRespMonthly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespMonthly struct{}"
	}

	return strings.Join([]string{"ListProductsRespMonthly", string(data)}, " ")
}
