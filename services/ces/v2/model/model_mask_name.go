package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaskName **参数解释**： 屏蔽规则名称。    **约束限制**： 不涉及。 **取值范围**： 只能为字母、数字、汉字、-、_，长度为[1,64]个字符。      **默认取值**： 不涉及。
type MaskName struct {
}

func (o MaskName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaskName struct{}"
	}

	return strings.Join([]string{"MaskName", string(data)}, " ")
}
