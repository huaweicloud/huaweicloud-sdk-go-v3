package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateId **参数解释**： 告警模板的ID。     **约束限制**： 不涉及。 **取值范围**： 以at开头，后跟字母、数字，长度为[2,64]个字符。           **默认取值**： 不涉及。
type TemplateId struct {
}

func (o TemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateId struct{}"
	}

	return strings.Join([]string{"TemplateId", string(data)}, " ")
}
