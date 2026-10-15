package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobName **参数解释**：Job名称。 **约束限制**：不涉及。 **取值范围**：长度为1~47的英文字母、数字、中划线的组合。 **默认取值**：不涉及。
type JobName struct {
}

func (o JobName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobName struct{}"
	}

	return strings.Join([]string{"JobName", string(data)}, " ")
}
