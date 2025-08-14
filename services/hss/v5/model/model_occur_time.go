package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OccurTime **参数解释**： 发生时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
type OccurTime struct {
}

func (o OccurTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OccurTime struct{}"
	}

	return strings.Join([]string{"OccurTime", string(data)}, " ")
}
