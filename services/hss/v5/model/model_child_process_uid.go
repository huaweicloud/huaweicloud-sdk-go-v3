package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChildProcessUid **参数解释**: 子进程用户id **取值范围**: 最小值0，最大值2147483647
type ChildProcessUid struct {
}

func (o ChildProcessUid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildProcessUid struct{}"
	}

	return strings.Join([]string{"ChildProcessUid", string(data)}, " ")
}
