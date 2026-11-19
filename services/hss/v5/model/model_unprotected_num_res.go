package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnprotectedNumRes **参数解释**: 未防护数 **取值范围**: 最小值0，最大值2147483647
type UnprotectedNumRes struct {
}

func (o UnprotectedNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnprotectedNumRes struct{}"
	}

	return strings.Join([]string{"UnprotectedNumRes", string(data)}, " ")
}
