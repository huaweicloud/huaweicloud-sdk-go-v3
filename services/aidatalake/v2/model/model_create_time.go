package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTime **参数解释**：创建时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type CreateTime struct {
}

func (o CreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTime struct{}"
	}

	return strings.Join([]string{"CreateTime", string(data)}, " ")
}
