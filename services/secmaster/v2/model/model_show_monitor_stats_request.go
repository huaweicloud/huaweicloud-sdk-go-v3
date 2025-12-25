package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorStatsRequest Request Object
type ShowMonitorStatsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *ShowPulsarMonitorStatsBody `json:"body,omitempty"`
}

func (o ShowMonitorStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitorStatsRequest", string(data)}, " ")
}
