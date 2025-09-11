package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTime **参数解释**： 屏蔽起始时间。          **约束限制**： 不涉及。 **取值范围**： 字符长度为8，格式为：HH:mm:ss         **默认取值**： 不涉及。
type StartTime struct {
}

func (o StartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTime struct{}"
	}

	return strings.Join([]string{"StartTime", string(data)}, " ")
}
