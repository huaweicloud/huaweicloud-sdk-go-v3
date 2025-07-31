package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessPid **参数解释**: 父进程id **取值范围**: 最小值0，最大值2147483647
type ParentProcessPid struct {
}

func (o ParentProcessPid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessPid struct{}"
	}

	return strings.Join([]string{"ParentProcessPid", string(data)}, " ")
}
