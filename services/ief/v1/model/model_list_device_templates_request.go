package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceTemplatesRequest Request Object
type ListDeviceTemplatesRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 终端设备名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认为1000
	Limit *string `json:"limit,omitempty"`
}

func (o ListDeviceTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListDeviceTemplatesRequest", string(data)}, " ")
}
