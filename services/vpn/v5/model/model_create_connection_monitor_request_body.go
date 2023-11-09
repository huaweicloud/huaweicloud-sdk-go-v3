package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectionMonitorRequestBody struct {
	ConnectionMonitor *CreateConnectionMonitorRequestBodyContent `json:"connection_monitor"`
}

func (o CreateConnectionMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectionMonitorRequestBody", string(data)}, " ")
}
