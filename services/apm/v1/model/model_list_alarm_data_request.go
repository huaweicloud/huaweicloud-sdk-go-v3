package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmDataRequest Request Object
type ListAlarmDataRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *AlarmDataListRequest `json:"body,omitempty"`
}

func (o ListAlarmDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDataRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmDataRequest", string(data)}, " ")
}
