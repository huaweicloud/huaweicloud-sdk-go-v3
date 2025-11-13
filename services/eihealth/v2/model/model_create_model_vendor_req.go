package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelVendorReq **参数解释**： 模型供应商修改请求体。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type CreateModelVendorReq struct {

	// **参数解释**： 模型供应商名称。 **约束限制**： 不涉及 **取值范围**： 支持中英文、数字、下划线、中划线、空格，长度2-64。 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 模型供应商英文名称。 **约束限制**： 不涉及 **取值范围**： 支持英文、数字、下划线、中划线、空格，长度2-64。 **默认取值**： 不涉及
	NameEn string `json:"name_en"`

	// **参数解释**： 模型供应商描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 鉴权类型。 **约束限制**： 不涉及 **取值范围**： * api_key：API Key **默认取值**： 不涉及
	AuthType string `json:"auth_type"`

	AuthInfo *VendorAuthInfo `json:"auth_info"`
}

func (o CreateModelVendorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelVendorReq struct{}"
	}

	return strings.Join([]string{"CreateModelVendorReq", string(data)}, " ")
}
