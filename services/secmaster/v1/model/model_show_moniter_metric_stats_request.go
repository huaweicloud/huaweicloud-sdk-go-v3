package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMoniterMetricStatsRequest Request Object
type ShowMoniterMetricStatsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ShowMetricStatsRequestBody `json:"body,omitempty"`
}

func (o ShowMoniterMetricStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMoniterMetricStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowMoniterMetricStatsRequest", string(data)}, " ")
}
