package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeNodesRequest struct {
	// 边缘节点名称，模糊匹配

	Name *string `json:"name,omitempty"`
	// 每页显示的条目数量，取值范围1~1000，默认为500

	Limit *string `json:"limit,omitempty"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *string `json:"offset,omitempty"`
	// 按终端设备ID查找

	DeviceId *string `json:"device_id,omitempty"`
	// 按绑定终端设备名称查找

	DeviceName *string `json:"device_name,omitempty"`
	// 按应用名称查找

	AppName *string `json:"app_name,omitempty"`
	// 标签的key和value通过点连接， 多个标签通过逗号连接，如：tags=key1.value1,key2.value2

	Tags *string `json:"tags,omitempty"`
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ListEdgeNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesRequest", string(data)}, " ")
}
