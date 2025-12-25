package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorStatsResponse Response Object
type ShowMonitorStatsResponse struct {
	EsResult *EsMonitorBody `json:"es_result,omitempty"`

	// pulsar返回的结果列表
	Result         *[]map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowMonitorStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorStatsResponse", string(data)}, " ")
}
