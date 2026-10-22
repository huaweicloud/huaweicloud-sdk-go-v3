package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTimeAdmin **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
type UpdateTimeAdmin struct {
}

func (o UpdateTimeAdmin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimeAdmin struct{}"
	}

	return strings.Join([]string{"UpdateTimeAdmin", string(data)}, " ")
}
