package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePropertiesRequest struct {
	// 下发属性的设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获得。

	DeviceId string `json:"device_id"`
	// Stage用户的Token, 仅提供给IoStage服务使用

	StageAuthToken *string `json:"Stage-Auth-Token,omitempty"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *DevicePropertiesRequest `json:"body,omitempty"`
}

func (o UpdatePropertiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePropertiesRequest struct{}"
	}

	return strings.Join([]string{"UpdatePropertiesRequest", string(data)}, " ")
}
