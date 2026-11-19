package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeRes **参数解释**: 时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
type TimeRes struct {
}

func (o TimeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRes struct{}"
	}

	return strings.Join([]string{"TimeRes", string(data)}, " ")
}
