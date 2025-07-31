package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChildProcessPid **参数解释**: 子进程id **取值范围**: 最小值0，最大值2147483647
type ChildProcessPid struct {
}

func (o ChildProcessPid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildProcessPid struct{}"
	}

	return strings.Join([]string{"ChildProcessPid", string(data)}, " ")
}
