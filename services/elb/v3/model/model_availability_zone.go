package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone 可用区。
type AvailabilityZone struct {

	// 参数解释：可用区唯一编码。
	Code string `json:"code"`

	// 参数解释：可用区状态。  取值范围：ACTIVE。
	State string `json:"state"`

	// [参数解释：未售罄的LB规格类别。  取值范围：L4 表示网络型LB未售罄；L7 表示应用型LB未售罄。](tag:hws,hws_hk,hws_eu,otc,tlf,ct,sbc,g42,hk_g42,dt_test,mix,hk_sbc,hws_ocb,dt) [参数解释：LB规格类别。  取值范围：L4 表示网络型LB；L7 表示应用型LB。](tag:ctc,cmcc,ocb,tm,srg,fcs,fcs_dt,hcso,hcso_dt,hk_vdf)
	Protocol []string `json:"protocol"`

	// 参数解释：网络公共边界组，如：center
	PublicBorderGroup string `json:"public_border_group"`

	// 参数解释：范围编码。  取值范围：0表示center，21表示homezone。
	Category int32 `json:"category"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
