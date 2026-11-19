package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeFlavorItemDetail struct {

	// **参数解释**: 规格ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 规格编码。 **取值范围**: 不涉及。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**: IaaS规格编码。 **取值范围**: 不涉及。
	IaasCode *string `json:"iaas_code,omitempty"`

	// **参数解释**: CPU核数。 **取值范围**: 不涉及。
	Vcpus *string `json:"vcpus,omitempty"`

	// **参数解释**: 内存大小，单位GB。 **取值范围**: 不涉及。
	Ram *string `json:"ram,omitempty"`

	// **参数解释**: 是否支持IPv6。 **取值范围**: - true：支持 - false：不支持
	SupportedIpv6 *bool `json:"supported_ipv6,omitempty"`

	// **参数解释**: 规格在各个可用区的状态。 **取值范围**: - normal：正常。 - sellout：已售罄。 - unknown：状态未知。 - unsupported：不支持。 - abandon：已废弃。
	AzStatus map[string]string `json:"az_status,omitempty"`
}

func (o ComputeFlavorItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavorItemDetail struct{}"
	}

	return strings.Join([]string{"ComputeFlavorItemDetail", string(data)}, " ")
}
