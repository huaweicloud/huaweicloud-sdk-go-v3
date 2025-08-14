package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessPid **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
type ProcessPid struct {
}

func (o ProcessPid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessPid struct{}"
	}

	return strings.Join([]string{"ProcessPid", string(data)}, " ")
}
