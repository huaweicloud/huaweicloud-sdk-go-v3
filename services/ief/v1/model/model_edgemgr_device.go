package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备属性
type EdgemgrDevice struct {

	// 终端设备ID，只允许英文字母、数字、下划线、中划线，必须以英文字母和数字开头，长度限制为24~64之间
	Id string `json:"id"`

	// 终端设备名称，只允许中文字符、英文字母、数字、下划线、中划线，长度限制为1~64
	Name string `json:"name"`

	// 访问协议，有如下选项： - userdefine：自定义协议 - modbus：modbus协议 - opc-ua：opc-ua协议 默认为userdefine
	AccessProtocol string `json:"access_protocol"`

	// 终端设备描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description string `json:"description"`

	// 静态属性
	Attributes map[string]ValueInAttributes `json:"attributes"`

	// 终端设备静态属性信息
	Twin map[string]ValueInTwinResponse `json:"twin"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 孪生属性配置
	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors"`

	Tags *ResourceTag `json:"tags"`
}

func (o EdgemgrDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDevice struct{}"
	}

	return strings.Join([]string{"EdgemgrDevice", string(data)}, " ")
}
