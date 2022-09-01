package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListConfigMapsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 配置项名称，模糊匹配
	Name *string `json:"name,omitempty" xml:"name"`

	// 每页显示的条目数量，取值范围1~1000，默认为1000
	Limit *string `json:"limit,omitempty" xml:"limit"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty" xml:"offset"`
}

func (o ListConfigMapsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigMapsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigMapsRequest", string(data)}, " ")
}
