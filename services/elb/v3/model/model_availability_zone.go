package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZone 可用区。
type AvailabilityZone struct {

	// 可用区唯一编码。
	Code string `json:"code"`

	// 可用区状态。  取值：ACTIVE。
	State string `json:"state"`

	// [当前可用区可用的LB规格类别列表。](tag:,hws_hk,ctc,hcs,cmcc,hws_ocb,tm,hws_eu,hcso_dt,dt,dt_test,hws_ocb,ocb,fcs,g42,hk_g42,hws_g42) [当前可用区未售罄的LB规格类别列表。](tag:hws)  取值：L4-表示当前可用区可创建网络型的LB；L7-表示当前可用区可创建应用型的LB。
	Protocol []string `json:"protocol"`

	// 可用区组，如：center
	PublicBorderGroup string `json:"public_border_group"`

	// 范围编码，0表示center，21表示homezone
	Category int32 `json:"category"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
