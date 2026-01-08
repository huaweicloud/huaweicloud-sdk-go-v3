package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone **参数解释**：可用区信息。
type AvailabilityZone struct {

	// **参数解释**：可用区唯一编码。  **取值范围**：不涉及
	Code string `json:"code"`

	// **参数解释**：可用区状态。  **取值范围**：ACTIVE。
	State string `json:"state"`

	// [**参数解释**：未售罄的LB规格类别。  **取值范围**：L4 表示网络型LB未售罄；L7 表示应用型LB未售罄。](tag:hws,hws_hk,hws_eu,otc,tlf,ct,sbc,g42,hk_g42,mix,hk_sbc,hws_ocb,dt) [**参数解释**：LB规格类别。  **取值范围**：L4 表示网络型LB；L7 表示应用型LB。](tag:ctc,cmcc,ocb,tm,srg,fcs,fcs_dt,hcso,hcso_dt,hk_vdf)
	Protocol []string `json:"protocol"`

	// **参数解释**：公网边界组。  **取值范围**： - center：表示中心站点的公网边界组 - 边缘站点名称：表示边缘站点的公网边界组
	PublicBorderGroup string `json:"public_border_group"`

	// **参数解释**：可用区子类型编码。该字段主要用于区分在边缘场景下，边缘AZ的类型。  **取值范围**：0表示center，21表示homezone，41表示IES。
	Category int32 `json:"category"`

	// **参数解释**：可用区的产品编码，用于控制台购买ELB前查询定价，仅边缘场景有效。  **取值范围**：不涉及 [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
