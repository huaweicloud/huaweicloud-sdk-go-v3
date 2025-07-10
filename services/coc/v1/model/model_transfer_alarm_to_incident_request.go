package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferAlarmToIncidentRequest Request Object
type TransferAlarmToIncidentRequest struct {
	Body *AlarmToIncidentRequestBody `json:"body,omitempty"`
}

func (o TransferAlarmToIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferAlarmToIncidentRequest struct{}"
	}

	return strings.Join([]string{"TransferAlarmToIncidentRequest", string(data)}, " ")
}
