package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessGid **参数解释**: 进程组ID **取值范围**: 最小值0，最大值2147483647
type ProcessGid struct {
}

func (o ProcessGid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessGid struct{}"
	}

	return strings.Join([]string{"ProcessGid", string(data)}, " ")
}
