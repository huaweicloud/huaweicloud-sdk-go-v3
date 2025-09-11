package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDate **参数解释**： 屏蔽起始日期。           **约束限制**： 不涉及。 **取值范围**： 字符长度为10，格式为：yyyy-MM-dd           **默认取值**： 不涉及。
type StartDate struct {
}

func (o StartDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDate struct{}"
	}

	return strings.Join([]string{"StartDate", string(data)}, " ")
}
