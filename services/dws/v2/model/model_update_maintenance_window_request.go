package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMaintenanceWindowRequest Request Object
type UpdateMaintenanceWindowRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	Body *MaintenanceWindow `json:"body,omitempty"`
}

func (o UpdateMaintenanceWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMaintenanceWindowRequest struct{}"
	}

	return strings.Join([]string{"UpdateMaintenanceWindowRequest", string(data)}, " ")
}
