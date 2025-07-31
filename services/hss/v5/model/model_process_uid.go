package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessUid **参数解释**： 进程名称 **取值范围**： 最小值0，最大值2147483647
type ProcessUid struct {
}

func (o ProcessUid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessUid struct{}"
	}

	return strings.Join([]string{"ProcessUid", string(data)}, " ")
}
