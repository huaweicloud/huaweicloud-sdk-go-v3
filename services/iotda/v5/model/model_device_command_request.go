package model

import (
	"encoding/json"

	"strings"
)

type DeviceCommandRequest struct {
	// 设备命令所属的设备服务ID，在设备关联的产品模型中定义。
	ServiceId *string `json:"service_id,omitempty"`
	// 设备命令名称，在设备关联的产品模型中定义。
	CommandName *string `json:"command_name,omitempty"`
	// 设备执行的命令，Json格式，里面是一个个健值对，如果serviceId不为空，每个健都是profile中命令的参数名（paraName）;如果serviceId为空则由用户自定义命令格式。设备命令示例：{\"value\":\"1\"}，具体格式需要应用和设备约定。
	Paras *interface{} `json:"paras"`
}

func (o DeviceCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeviceCommandRequest struct{}"
	}

	return strings.Join([]string{"DeviceCommandRequest", string(data)}, " ")
}
