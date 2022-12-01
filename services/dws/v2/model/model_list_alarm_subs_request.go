package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmSubsRequest struct {
}

func (o ListAlarmSubsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmSubsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmSubsRequest", string(data)}, " ")
}
