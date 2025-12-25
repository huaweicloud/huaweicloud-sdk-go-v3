package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Level **参数解释**： 告警级别。    **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。 - 1：紧急 - 2：重要 - 3：次要 - 4：提示         **默认取值**： 2
type Level struct {
}

func (o Level) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Level struct{}"
	}

	return strings.Join([]string{"Level", string(data)}, " ")
}
