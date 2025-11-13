package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VendorAuthInfo **参数解释**： 鉴权信息。 **约束限制**： 当鉴权类型为api_key时，api_key必填。 **取值范围**： 不涉及 **默认取值**： 不涉及
type VendorAuthInfo struct {

	// **参数解释**： ak。 **约束限制**： 不涉及 **取值范围**： 字符长度为[1-1024]。 **默认取值**： 不涉及
	Ak *string `json:"ak,omitempty"`

	// **参数解释**： sk。 **约束限制**： 不涉及 **取值范围**： 字符长度为[1-1024]。 **默认取值**： 不涉及
	Sk *string `json:"sk,omitempty"`

	// **参数解释**： api_key。 **约束限制**： 不涉及 **取值范围**： 字符长度为[1-1024]。 **默认取值**： 不涉及
	ApiKey *string `json:"api_key,omitempty"`
}

func (o VendorAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorAuthInfo struct{}"
	}

	return strings.Join([]string{"VendorAuthInfo", string(data)}, " ")
}
