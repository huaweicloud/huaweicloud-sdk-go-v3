package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainInfo struct {

	// **参数解释**： 域名地址id **取值范围**： 不涉及
	DomainAddressId *string `json:"domain_address_id,omitempty"`

	// **参数解释**： 域名 **取值范围**： 不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**： 域名描述 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o DomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainInfo struct{}"
	}

	return strings.Join([]string{"DomainInfo", string(data)}, " ")
}
