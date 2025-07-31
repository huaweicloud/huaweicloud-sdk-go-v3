package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessEgid **参数解释**: 父进程有效组ID **取值范围**: 最小值0，最大值2147483647
type ParentProcessEgid struct {
}

func (o ParentProcessEgid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessEgid struct{}"
	}

	return strings.Join([]string{"ParentProcessEgid", string(data)}, " ")
}
