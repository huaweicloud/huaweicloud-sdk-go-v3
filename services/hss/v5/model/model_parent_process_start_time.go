package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParentProcessStartTime **参数解释**: 父进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
type ParentProcessStartTime struct {
}

func (o ParentProcessStartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParentProcessStartTime struct{}"
	}

	return strings.Join([]string{"ParentProcessStartTime", string(data)}, " ")
}
