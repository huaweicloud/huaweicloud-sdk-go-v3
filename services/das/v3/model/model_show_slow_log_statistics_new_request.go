package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogStatisticsNewRequest Request Object
type ShowSlowLogStatisticsNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ShowSlowLogStatisticsNewRequestBody `json:"body,omitempty"`
}

func (o ShowSlowLogStatisticsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsNewRequest", string(data)}, " ")
}
