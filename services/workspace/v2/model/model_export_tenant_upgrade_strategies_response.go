package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTenantUpgradeStrategiesResponse Response Object
type ExportTenantUpgradeStrategiesResponse struct {

	// 导出任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportTenantUpgradeStrategiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTenantUpgradeStrategiesResponse struct{}"
	}

	return strings.Join([]string{"ExportTenantUpgradeStrategiesResponse", string(data)}, " ")
}
