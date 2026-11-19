package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BeginTimeRes **参数解释**： 结束时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
type BeginTimeRes struct {
}

func (o BeginTimeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BeginTimeRes struct{}"
	}

	return strings.Join([]string{"BeginTimeRes", string(data)}, " ")
}
