package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobId **参数解释**：作业ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
type JobId struct {
}

func (o JobId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobId struct{}"
	}

	return strings.Join([]string{"JobId", string(data)}, " ")
}
