package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferAlarmToIncidentResponse Response Object
type TransferAlarmToIncidentResponse struct {

	// 返回数据
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o TransferAlarmToIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferAlarmToIncidentResponse struct{}"
	}

	return strings.Join([]string{"TransferAlarmToIncidentResponse", string(data)}, " ")
}
