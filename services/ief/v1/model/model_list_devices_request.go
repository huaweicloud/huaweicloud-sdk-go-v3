package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDevicesRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 终端设备名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 节点ID, 精确匹配
	NodeId *string `json:"node_id,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认为1000
	Limit *string `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty"`

	// 是否绑定到边缘节点，为“true”时返回所有已绑定到节点的设备列表，为“false”则返回未绑定节点的设备列表
	IsBinding *string `json:"is_binding,omitempty"`

	// 标签的key和value通过点连接， 多个标签通过逗号连接，如：tags=key1.value1,key2.value2
	Tags *string `json:"tags,omitempty"`
}

func (o ListDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}
