package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPropertiesRequest struct {
	// 下发属性的设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获得。

	DeviceId string `json:"device_id"`
	// Stage用户的Token, 仅提供给IoStage服务使用

	StageAuthToken *string `json:"Stage-Auth-Token,omitempty"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 设备的服务ID，在设备关联的产品模型中定义。

	ServiceId string `json:"service_id"`
}

func (o ListPropertiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListPropertiesRequest", string(data)}, " ")
}
