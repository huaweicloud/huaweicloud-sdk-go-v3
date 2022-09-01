package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点参数
type EdgeNode struct {

	// 边缘节点名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64 Name为必填字段，且本帐号中唯一。
	Name string `json:"name" xml:"name"`

	// 边缘节点描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty" xml:"description"`

	// 边缘节点是否开启GPU，默认为false
	EnableGpu *bool `json:"enable_gpu,omitempty" xml:"enable_gpu"`

	// 边缘节点日志配置
	LogConfigs *[]LogConfigs `json:"log_configs,omitempty" xml:"log_configs"`

	// 关联设备信息
	DeviceInfos *[]DeviceInfos `json:"device_infos,omitempty" xml:"device_infos"`

	// 边缘节点是否开启NPU，true表示开启，false表示不开启，默认为false
	EnableNpu *bool `json:"enable_npu,omitempty" xml:"enable_npu"`

	// NPU类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty" xml:"npu_type"`

	// 边缘节点属性，关联属性个数最多为32个
	Attributes *[]Attributes `json:"attributes,omitempty" xml:"attributes"`

	// 边缘节点启用Docker，默认为true
	EnableDocker *bool `json:"enable_docker,omitempty" xml:"enable_docker"`

	// 边缘节点标签，标签个数最多为20个
	Tags *[]NodeResTag `json:"tags,omitempty" xml:"tags"`

	MqttConfig *MqttConfigs `json:"mqtt_config,omitempty" xml:"mqtt_config"`
}

func (o EdgeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNode struct{}"
	}

	return strings.Join([]string{"EdgeNode", string(data)}, " ")
}
