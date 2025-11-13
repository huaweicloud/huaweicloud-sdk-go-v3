package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelVendorReq **参数解释**： 模型供应商修改请求体。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type UpdateModelVendorReq struct {

	// **参数解释**： 模型供应商名称。 **约束限制**： 不涉及 **取值范围**： 支持中英文、数字、下划线、中划线、空格，长度2-64。 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模型供应商英文名称。 **约束限制**： 不涉及 **取值范围**： 支持英文、数字、下划线、中划线、空格，长度2-64。 **默认取值**： 不涉及
	NameEn *string `json:"name_en,omitempty"`

	// **参数解释**： 模型供应商描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 鉴权类型。 **约束限制**： 不涉及 **取值范围**： * api_key：API Key **默认取值**： 不涉及
	AuthType *string `json:"auth_type,omitempty"`

	AuthInfo *VendorAuthInfo `json:"auth_info,omitempty"`
}

func (o UpdateModelVendorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelVendorReq struct{}"
	}

	return strings.Join([]string{"UpdateModelVendorReq", string(data)}, " ")
}
