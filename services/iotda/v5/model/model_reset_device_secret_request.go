package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetDeviceSecretRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 设备ID，用于唯一标识一个设备。在注册设备时直接指定，或者由物联网平台分配获得。由物联网平台分配时，生成规则为\"product_id\" + \"_\" + \"node_id\"拼接而成。

	DeviceId string `json:"device_id"`
	// 对设备执行的操作，取值范围：resetSecret。 - resetSecret: 重置密钥。注意：NB设备密钥由于协议特殊性，只支持十六进制密钥接入。

	ActionId string `json:"action_id"`

	Body *ResetDeviceSecret `json:"body,omitempty"`
}

func (o ResetDeviceSecretRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetDeviceSecretRequest struct{}"
	}

	return strings.Join([]string{"ResetDeviceSecretRequest", string(data)}, " ")
}
