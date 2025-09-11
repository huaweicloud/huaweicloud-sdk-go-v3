package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceVersionRequest Request Object
type UpgradeInstanceVersionRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *OpenGaussUpgradeRequest `json:"body,omitempty"`
}

func (o UpgradeInstanceVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceVersionRequest", string(data)}, " ")
}
