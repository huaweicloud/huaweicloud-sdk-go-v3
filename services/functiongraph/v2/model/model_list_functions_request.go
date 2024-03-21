package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionsRequest Request Object
type ListFunctionsRequest struct {

	// 上一次查询到的最后的记录位置。
	Marker *string `json:"marker,omitempty"`

	// 每次查询获取的最大函数记录数量 最大值：400 如果不提供该值或者提供的值大于400或等于0，则使用默认值：400 如果该值小于0，则返回参数错误。
	Maxitems *string `json:"maxitems,omitempty"`

	// 自定义分组名称。
	PackageName *string `json:"package_name,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionsRequest", string(data)}, " ")
}
