package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivateNetworkSegmentVo 私网网段的详细信息，会将参数中conf_id的私网网段信息覆盖更新
type PrivateNetworkSegmentVo struct {

	// **参数解释**： 私网网段ID，更新私网网段时需要 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ConfId *string `json:"conf_id,omitempty"`

	// **参数解释**： 私网网段，格式为：IP/Mask **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ProtectNetwork string `json:"protect_network"`

	// **参数解释**： 私网网段的可用区信息，如AZ1 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AzInfo *string `json:"az_info,omitempty"`

	// **参数解释**： 子网名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SubnetName *string `json:"subnet_name,omitempty"`
}

func (o PrivateNetworkSegmentVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateNetworkSegmentVo struct{}"
	}

	return strings.Join([]string{"PrivateNetworkSegmentVo", string(data)}, " ")
}
