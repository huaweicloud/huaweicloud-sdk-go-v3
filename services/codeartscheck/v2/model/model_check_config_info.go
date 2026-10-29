package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckConfigInfo **参数解释：** 检查参数配置信息。
type CheckConfigInfo struct {

	// **参数解释：** 检查参数名称。 **取值范围:** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 检查参数对应的Key值。 **取值范围:** 不涉及。
	CfgKey *string `json:"cfg_key,omitempty"`

	// **参数解释：** 检查参数默认值。 **取值范围:** 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释：** 检查参数可选项。 **取值范围:** 不涉及。
	OptionValue *string `json:"option_value,omitempty"`

	// **参数解释：** 是否必填参数。 **取值范围:** - 0：非必填。 - 1：必填。
	IsRequired *int32 `json:"is_required,omitempty"`

	// **参数解释：** 检查参数说明。 **取值范围:** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 参数类型。 **取值范围:** - 0：文本。 - 2：有可选项。
	Type *int32 `json:"type,omitempty"`

	// **参数解释：** 参数状态。 **取值范围:** - on：启用。 - off：未启用。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 检查参数值。 **取值范围:** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o CheckConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConfigInfo struct{}"
	}

	return strings.Join([]string{"CheckConfigInfo", string(data)}, " ")
}
