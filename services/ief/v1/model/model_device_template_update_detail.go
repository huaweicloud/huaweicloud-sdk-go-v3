package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备模板
type DeviceTemplateUpdateDetail struct {

	// 设备模板描述，最大长度255
	Description *string `json:"description,omitempty" xml:"description"`

	// 终端设备静态属性，最多64个键值。
	Attributes map[string]ValueInAttributes `json:"attributes,omitempty" xml:"attributes"`

	// 终端设备动态属性
	Twin map[string]ValueInTwin `json:"twin,omitempty" xml:"twin"`

	Tags *DeviceTemplateUpdateDetailTags `json:"tags,omitempty" xml:"tags"`

	// - userdefine：自定义协议 - modbus：modbus协议 - opc-ua：opc-ua协议
	AccessProtocol *string `json:"access_protocol,omitempty" xml:"access_protocol"`

	// 孪生属性配置，与access_protocol关联。
	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors,omitempty" xml:"property_visitors"`
}

func (o DeviceTemplateUpdateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceTemplateUpdateDetail struct{}"
	}

	return strings.Join([]string{"DeviceTemplateUpdateDetail", string(data)}, " ")
}
