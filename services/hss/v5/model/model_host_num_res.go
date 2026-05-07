package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostNumRes **参数解释**: 防护主机数量。 **取值范围**: 最小值1，最大值2000000
type HostNumRes struct {
}

func (o HostNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNumRes struct{}"
	}

	return strings.Join([]string{"HostNumRes", string(data)}, " ")
}
