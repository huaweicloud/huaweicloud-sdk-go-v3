package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Properties 资产详细属性。
type Properties struct {
	HwcEcs *HwcEcs `json:"hwc_ecs,omitempty"`

	HwcEip *HwcEip `json:"hwc_eip,omitempty"`

	HwcVpc *HwcVpc `json:"hwc_vpc,omitempty"`

	HwcSubnet *HwcSubnet `json:"hwc_subnet,omitempty"`

	HwcRds *HwcRds `json:"hwc_rds,omitempty"`

	HwcDomain *HwcDomain `json:"hwc_domain,omitempty"`

	Website *Website `json:"website,omitempty"`

	OcaIp *OcaIp `json:"oca_ip,omitempty"`
}

func (o Properties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Properties struct{}"
	}

	return strings.Join([]string{"Properties", string(data)}, " ")
}
