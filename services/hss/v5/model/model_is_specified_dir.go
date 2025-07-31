package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsSpecifiedDir **参数解释**： 是否指定学习目录 **约束限制**： 不涉及 **取值范围**: - true：是 - false：否 **默认取值**： 不涉及
type IsSpecifiedDir struct {
}

func (o IsSpecifiedDir) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsSpecifiedDir struct{}"
	}

	return strings.Join([]string{"IsSpecifiedDir", string(data)}, " ")
}
