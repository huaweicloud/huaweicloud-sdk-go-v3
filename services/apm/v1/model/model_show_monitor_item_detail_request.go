package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorItemDetailRequest Request Object
type ShowMonitorItemDetailRequest struct {

	// 监控项ID
	MonitorItemId int64 `json:"monitor_item_id"`

	// 环境ID
	EnvId int64 `json:"env_id"`

	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowMonitorItemDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorItemDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitorItemDetailRequest", string(data)}, " ")
}
