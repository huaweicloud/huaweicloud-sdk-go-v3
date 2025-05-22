package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrowCheckResult **参数解释**： 扩容前检查响应体。 当请求header中x-language字段为zh_cn时返回中文描述、类型信息。 当请求header中x-language字段为en_us时返回英文描述、类型信息。 **取值范围**： 不涉及。
type GrowCheckResult struct {

	// **参数解释**： 检查是否通过，检查通过项默认不展示。 **取值范围**： true/false
	Pass *bool `json:"pass,omitempty"`

	// **参数解释**： 检查不通过的原因描述。 **取值范围**： 不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： 是否必须检查项。 **取值范围**： - true：必须，校验不通过时不允许扩容，继续扩容也会失败 - false：非必须，校验不通过时允许扩容，仅做提示告知风险
	Required *bool `json:"required,omitempty"`

	// **参数解释**： 描述信息。 **取值范围**： 不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 分类。 **取值范围**： 配额、权限、版本、状态
	Type *string `json:"type,omitempty"`
}

func (o GrowCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrowCheckResult struct{}"
	}

	return strings.Join([]string{"GrowCheckResult", string(data)}, " ")
}
