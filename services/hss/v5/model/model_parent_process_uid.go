package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessUid **参数解释**: 父进程用户id **取值范围**: 最小值0，最大值2147483647
type ParentProcessUid struct {
}

func (o ParentProcessUid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessUid struct{}"
	}

	return strings.Join([]string{"ParentProcessUid", string(data)}, " ")
}
