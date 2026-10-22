package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTime **参数解释**：资源更新时间，使用UTC时间，ISO 8601格式，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
type UpdateTime struct {
}

func (o UpdateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTime struct{}"
	}

	return strings.Join([]string{"UpdateTime", string(data)}, " ")
}
