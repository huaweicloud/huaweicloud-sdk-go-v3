package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceTemplateUpdateDetail 设备模板
type DeviceTemplateUpdateDetail struct {

	// 设备模板描述，最大长度255
	Description *string `json:"description,omitempty"`

	// 终端设备静态属性，最多64个键值。
	Attributes map[string]ValueInAttributes `json:"attributes,omitempty"`

	// 终端设备动态属性
	Twin map[string]ValueInTwin `json:"twin,omitempty"`

	Tags *DeviceTemplateUpdateDetailTags `json:"tags,omitempty"`

	// - userdefine：自定义协议 - modbus：modbus协议 - opc-ua：opc-ua协议
	AccessProtocol *string `json:"access_protocol,omitempty"`

	// 孪生属性配置，与access_protocol关联。
	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors,omitempty"`
}

func (o DeviceTemplateUpdateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceTemplateUpdateDetail struct{}"
	}

	return strings.Join([]string{"DeviceTemplateUpdateDetail", string(data)}, " ")
}
