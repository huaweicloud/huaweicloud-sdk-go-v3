package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceMaintenanceWindowResponse Response Object
type ModifyInstanceMaintenanceWindowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyInstanceMaintenanceWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceMaintenanceWindowResponse struct{}"
	}

	return strings.Join([]string{"ModifyInstanceMaintenanceWindowResponse", string(data)}, " ")
}
