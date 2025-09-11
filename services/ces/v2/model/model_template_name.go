package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateName **参数解释**： 告警模板的名称。 **约束限制**： 不涉及。 **取值范围**： 以字母或汉字开头，可包含字母、数字、汉字、_、-，长度为[1,128]个字符。           **默认取值**： 不涉及。
type TemplateName struct {
}

func (o TemplateName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateName struct{}"
	}

	return strings.Join([]string{"TemplateName", string(data)}, " ")
}
