package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTimeAdmin **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
type CreateTimeAdmin struct {
}

func (o CreateTimeAdmin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTimeAdmin struct{}"
	}

	return strings.Join([]string{"CreateTimeAdmin", string(data)}, " ")
}
