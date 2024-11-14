package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecondLevelMonitoringStatusResponse Response Object
type ShowSecondLevelMonitoringStatusResponse struct {

	// 秒级监控开启状态。 - true: 开启; - false: 关闭。
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
