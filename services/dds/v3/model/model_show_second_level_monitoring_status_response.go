package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecondLevelMonitoringStatusResponse Response Object
type ShowSecondLevelMonitoringStatusResponse struct {

	// 秒级监控开启状态。  取值为true,开启，取值为false，关闭。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSecondLevelMonitoringStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecondLevelMonitoringStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSecondLevelMonitoringStatusResponse", string(data)}, " ")
}
