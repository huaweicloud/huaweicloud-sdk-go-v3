package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmNotifyRequest struct {

	// 应用id，用于鉴权。
	XBusinessId int64 `json:"x-business-id"`

	Body *AlarmNotifyListRequest `json:"body,omitempty"`
}

func (o ListAlarmNotifyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmNotifyRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmNotifyRequest", string(data)}, " ")
}
