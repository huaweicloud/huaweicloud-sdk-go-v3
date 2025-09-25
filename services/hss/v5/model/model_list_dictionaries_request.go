package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDictionariesRequest Request Object
type ListDictionariesRequest struct {

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 使用场景（如果没有区分可以忽略） **约束限制**: 不涉及 **取值范围**: - hws： 国内站。 - hec-hk：国际站。  **默认取值**: 不涉及
	Scene *string `json:"scene,omitempty"`

	// **参数解释**: 字典分组 **约束限制**: 不涉及 **取值范围**: - featureSwitch： 页面特性开关。  **默认取值**: 不涉及
	GroupCode string `json:"group_code"`

	// **参数解释**: 字典项编码，group_code下字典项编码不重复 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位  **默认取值**: 不涉及
	Code *string `json:"code,omitempty"`
}

func (o ListDictionariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDictionariesRequest struct{}"
	}

	return strings.Join([]string{"ListDictionariesRequest", string(data)}, " ")
}
