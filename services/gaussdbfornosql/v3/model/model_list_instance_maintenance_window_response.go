package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceMaintenanceWindowResponse Response Object
type ListInstanceMaintenanceWindowResponse struct {

	// 参数解释： 实例的可维护时间窗。
	MaintenanceWindow *string `json:"maintenance_window,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ListInstanceMaintenanceWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMaintenanceWindowResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceMaintenanceWindowResponse", string(data)}, " ")
}
