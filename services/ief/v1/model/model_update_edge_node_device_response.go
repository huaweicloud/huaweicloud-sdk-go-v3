package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeNodeDeviceResponse struct {
	// 工业终端设备预留字段

	DeleteConnector *bool `json:"delete_connector,omitempty"`
	// 工业终端设备预留字段

	DeployConnector *bool `json:"deploy_connector,omitempty"`
	// 工业终端设备预留字段

	DeploymentId *string `json:"deployment_id,omitempty"`

	UpdateDevices  *NodeDevice `json:"update_devices,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateEdgeNodeDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeDeviceResponse", string(data)}, " ")
}
