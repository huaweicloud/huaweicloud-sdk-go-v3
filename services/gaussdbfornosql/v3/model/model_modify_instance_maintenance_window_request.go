package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceMaintenanceWindowRequest Request Object
type ModifyInstanceMaintenanceWindowRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyInstanceOpsWindowV3Req `json:"body,omitempty"`
}

func (o ModifyInstanceMaintenanceWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceMaintenanceWindowRequest struct{}"
	}

	return strings.Join([]string{"ModifyInstanceMaintenanceWindowRequest", string(data)}, " ")
}
