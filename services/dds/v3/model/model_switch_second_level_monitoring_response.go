package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchSecondLevelMonitoringResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSecondLevelMonitoringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringResponse struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringResponse", string(data)}, " ")
}
