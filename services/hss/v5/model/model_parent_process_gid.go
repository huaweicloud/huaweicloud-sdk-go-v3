package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessGid **参数解释**: 父进程组ID **取值范围**: 最小值0，最大值2147483647
type ParentProcessGid struct {
}

func (o ParentProcessGid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessGid struct{}"
	}

	return strings.Join([]string{"ParentProcessGid", string(data)}, " ")
}
