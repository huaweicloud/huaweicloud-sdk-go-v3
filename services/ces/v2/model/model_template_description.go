package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateDescription **参数解释**： 告警模板的描述     **约束限制**： 不涉及。 **取值范围**： 长度范围[0,256]。          **默认取值**： 空字符串。
type TemplateDescription struct {
}

func (o TemplateDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateDescription struct{}"
	}

	return strings.Join([]string{"TemplateDescription", string(data)}, " ")
}
