package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTimeRes **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
type UpdateTimeRes struct {
}

func (o UpdateTimeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimeRes struct{}"
	}

	return strings.Join([]string{"UpdateTimeRes", string(data)}, " ")
}
