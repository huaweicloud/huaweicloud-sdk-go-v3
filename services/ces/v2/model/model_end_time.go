package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndTime **参数解释**： 屏蔽截止时间。          **约束限制**： 不涉及。 **取值范围**： 字符长度为8，格式为：HH:mm:ss         **默认取值**： 不涉及。
type EndTime struct {
}

func (o EndTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndTime struct{}"
	}

	return strings.Join([]string{"EndTime", string(data)}, " ")
}
