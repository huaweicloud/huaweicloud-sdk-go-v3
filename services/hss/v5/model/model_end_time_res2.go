package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndTimeRes2 **参数解释**： 结束时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
type EndTimeRes2 struct {
}

func (o EndTimeRes2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndTimeRes2 struct{}"
	}

	return strings.Join([]string{"EndTimeRes2", string(data)}, " ")
}
