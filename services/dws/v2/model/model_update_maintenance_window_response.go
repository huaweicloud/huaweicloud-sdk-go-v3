package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMaintenanceWindowResponse Response Object
type UpdateMaintenanceWindowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMaintenanceWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMaintenanceWindowResponse struct{}"
	}

	return strings.Join([]string{"UpdateMaintenanceWindowResponse", string(data)}, " ")
}
