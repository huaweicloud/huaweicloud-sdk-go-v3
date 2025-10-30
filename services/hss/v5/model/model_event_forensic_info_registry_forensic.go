package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoRegistryForensic **参数解释**： 注册表取证信息 **取值范围**： 不涉及
type EventForensicInfoRegistryForensic struct {

	// **参数解释**： 注册表KEY **取值范围**： 不涉及
	RegKey *string `json:"reg_key,omitempty"`

	// **参数解释**： 注册表VALUE **取值范围**： 不涉及
	RegValue *string `json:"reg_value,omitempty"`

	// **参数解释**： 注册表新KEY **取值范围**： 不涉及
	RegNewKey *string `json:"reg_new_key,omitempty"`

	// **参数解释**： 注册表KEY/VALUE操作类型 **取值范围**： 不涉及
	RegOpType *string `json:"reg_op_type,omitempty"`
}

func (o EventForensicInfoRegistryForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoRegistryForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoRegistryForensic", string(data)}, " ")
}
