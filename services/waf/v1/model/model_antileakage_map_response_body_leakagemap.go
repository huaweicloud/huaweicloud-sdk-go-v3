package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntileakageMapResponseBodyLeakagemap **参数解释：** 敏感信息的内容类型和返回码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type AntileakageMapResponseBodyLeakagemap struct {

	// **参数解释：** 敏感信息的内容类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sensitive *[]string `json:"sensitive,omitempty"`

	// **参数解释：** 敏感信息的返回码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Code *[]string `json:"code,omitempty"`
}

func (o AntileakageMapResponseBodyLeakagemap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntileakageMapResponseBodyLeakagemap struct{}"
	}

	return strings.Join([]string{"AntileakageMapResponseBodyLeakagemap", string(data)}, " ")
}
