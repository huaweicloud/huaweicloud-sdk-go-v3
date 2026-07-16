package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolMonitorResponse Response Object
type ShowPoolMonitorResponse struct {

	// **参数解释**：资源池的监控指标数据。
	Metrics        *[]PoolMonitorMetrics `json:"metrics,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowPoolMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolMonitorResponse", string(data)}, " ")
}
