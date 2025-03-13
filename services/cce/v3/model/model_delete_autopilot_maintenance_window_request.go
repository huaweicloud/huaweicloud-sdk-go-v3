package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutopilotMaintenanceWindowRequest Request Object
type DeleteAutopilotMaintenanceWindowRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o DeleteAutopilotMaintenanceWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutopilotMaintenanceWindowRequest struct{}"
	}

	return strings.Join([]string{"DeleteAutopilotMaintenanceWindowRequest", string(data)}, " ")
}
