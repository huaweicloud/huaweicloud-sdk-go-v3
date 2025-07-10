package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsRequest Request Object
type ListAlarmsRequest struct {

	// 告警级别 | 1 - 紧急 2 - 重要 3 - 次要 4 - 提示。
	Level *int32 `json:"level,omitempty"`

	// 分页偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
