package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorItemViewConfigRequest Request Object
type ShowMonitorItemViewConfigRequest struct {

	// 环境id。
	EnvId int64 `json:"env_id"`

	// 采集器id。
	CollectorId int64 `json:"collector_id"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowMonitorItemViewConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorItemViewConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowMonitorItemViewConfigRequest", string(data)}, " ")
}
