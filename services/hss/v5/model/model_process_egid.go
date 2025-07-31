package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessEgid **参数解释**: 进程有效组ID **取值范围**: 最小值0，最大值2147483647
type ProcessEgid struct {
}

func (o ProcessEgid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessEgid struct{}"
	}

	return strings.Join([]string{"ProcessEgid", string(data)}, " ")
}
