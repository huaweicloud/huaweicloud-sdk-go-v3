package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationBasicInfo 应用基本信息
type ApplicationBasicInfo struct {

	// **参数解释**： 应用名称。 **约束限制**： 不涉及。 **取值范围**： [1-200]个字符。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 应用图标URL，当前仅支持预置应用图标，取值为0~9的数字，分别标识第1到第10张预置图标。 **约束限制**： 0~9的数字。 **取值范围**： [1-200]的字符串。 **默认取值**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 应用描述。 **约束限制**： 不涉及。 **取值范围**： [1-200]个字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ApplicationBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationBasicInfo struct{}"
	}

	return strings.Join([]string{"ApplicationBasicInfo", string(data)}, " ")
}
