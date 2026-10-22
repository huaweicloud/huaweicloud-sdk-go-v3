package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTime **参数解释**：资源创建时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
type CreateTime struct {
}

func (o CreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTime struct{}"
	}

	return strings.Join([]string{"CreateTime", string(data)}, " ")
}
