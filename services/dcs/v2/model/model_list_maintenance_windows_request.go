package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMaintenanceWindowsRequest struct {
}

func (o ListMaintenanceWindowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMaintenanceWindowsRequest struct{}"
	}

	return strings.Join([]string{"ListMaintenanceWindowsRequest", string(data)}, " ")
}
