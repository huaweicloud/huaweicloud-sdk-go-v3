package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssistantModelVendorInfo 供应商基本信息。
type AssistantModelVendorInfo struct {

	// **参数解释**： 模型供应商名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	VendorName *string `json:"vendor_name,omitempty"`

	// **参数解释**： 模型供应商类型。 **约束限制**： 不涉及 **取值范围**： * SYSTEM：预置供应商 * CUSTOM：自定义供应商 **默认取值**： 不涉及
	VendorType *string `json:"vendor_type,omitempty"`

	// **参数解释**： 供应商模型列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Models *[]ModelInfo `json:"models,omitempty"`
}

func (o AssistantModelVendorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssistantModelVendorInfo struct{}"
	}

	return strings.Join([]string{"AssistantModelVendorInfo", string(data)}, " ")
}
