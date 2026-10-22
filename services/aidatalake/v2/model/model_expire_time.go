package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpireTime **参数解释**：超期时间，null表示永久有效。使用UTC时间，ISO 8601格式，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
type ExpireTime struct {
}

func (o ExpireTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpireTime struct{}"
	}

	return strings.Join([]string{"ExpireTime", string(data)}, " ")
}
