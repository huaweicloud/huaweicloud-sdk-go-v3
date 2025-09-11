package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndDate **参数解释**： 屏蔽截止日期。           **约束限制**： 不涉及。 **取值范围**： 字符长度为10，格式为：yyyy-MM-dd           **默认取值**： 不涉及。
type EndDate struct {
}

func (o EndDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndDate struct{}"
	}

	return strings.Join([]string{"EndDate", string(data)}, " ")
}
