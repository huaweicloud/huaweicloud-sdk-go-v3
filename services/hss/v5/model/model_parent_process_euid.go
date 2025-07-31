package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessEuid **参数解释**: 父进程有效用户ID **取值范围**: 最小值0，最大值2147483647
type ParentProcessEuid struct {
}

func (o ParentProcessEuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessEuid struct{}"
	}

	return strings.Join([]string{"ParentProcessEuid", string(data)}, " ")
}
