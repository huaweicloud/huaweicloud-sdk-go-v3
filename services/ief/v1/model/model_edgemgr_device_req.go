package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备属性
type EdgemgrDeviceReq struct {
	// 终端设备名称，只允许中文字符、英文字母、数字、下划线、中划线，长度限制为1~64

	Name string `json:"name"`
	// 访问协议，有如下选项： - userdefine：自定义协议 - modbus：modbus协议 - opc-ua：opc-ua协议 默认为userdefine

	AccessProtocol *string `json:"access_protocol,omitempty"`
	// 终端设备描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Description *string `json:"description,omitempty"`
	// 静态属性

	Attributes map[string]ValueInAttributes `json:"attributes,omitempty"`
	// 终端设备动态属性

	Twin map[string]ValueInTwin `json:"twin,omitempty"`

	Tags *ResourceTag `json:"tags,omitempty"`
	// 孪生属性配置

	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors,omitempty"`
}

func (o EdgemgrDeviceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDeviceReq struct{}"
	}

	return strings.Join([]string{"EdgemgrDeviceReq", string(data)}, " ")
}
