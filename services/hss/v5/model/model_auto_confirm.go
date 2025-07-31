package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoConfirm **参数解释**： 是否自动确认学习结果 **取值范围**: - true：是 - false：否
type AutoConfirm struct {
}

func (o AutoConfirm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoConfirm struct{}"
	}

	return strings.Join([]string{"AutoConfirm", string(data)}, " ")
}
