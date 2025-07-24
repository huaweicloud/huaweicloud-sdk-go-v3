package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmSummaryRequest Request Object
type ShowAlarmSummaryRequest struct {
}

func (o ShowAlarmSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmSummaryRequest", string(data)}, " ")
}
