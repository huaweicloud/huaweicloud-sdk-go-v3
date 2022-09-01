package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeNodeDeviceResponse struct {

	// 工业终端设备预留字段
	DeleteConnector *bool `json:"delete_connector,omitempty" xml:"delete_connector"`

	// 工业终端设备预留字段
	DeployConnector *bool `json:"deploy_connector,omitempty" xml:"deploy_connector"`

	// 工业终端设备预留字段
	DeploymentId *string `json:"deployment_id,omitempty" xml:"deployment_id"`

	UpdateDevices  *NodeDevice `json:"update_devices,omitempty" xml:"update_devices"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateEdgeNodeDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeDeviceResponse", string(data)}, " ")
}
