package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelVendor 模型供应商详情。
type ModelVendor struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 模型供应商名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模型供应商英文名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NameEn *string `json:"name_en,omitempty"`

	// **参数解释**： 模型供应商类型。 **约束限制**： 不涉及 **取值范围**： * SYSTEM：预置供应商 * CUSTOM：自定义供应商 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 模型服务状态。 **约束限制**： 不涉及 **取值范围**： * AVAILABLE：已接入 * UNAVAILABLE：未接入 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 模型供应商创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 模型供应商修改时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 模型供应商描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 鉴权类型。 **约束限制**： 不涉及 **取值范围**： * API_KEY：API Key **默认取值**： 不涉及
	AuthType *string `json:"auth_type,omitempty"`

	AuthInfo *VendorAuthInfo `json:"auth_info,omitempty"`

	// **参数解释**： 内置供应商获取鉴权信息地址。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AuthUrl *string `json:"auth_url,omitempty"`
}

func (o ModelVendor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelVendor struct{}"
	}

	return strings.Join([]string{"ModelVendor", string(data)}, " ")
}
