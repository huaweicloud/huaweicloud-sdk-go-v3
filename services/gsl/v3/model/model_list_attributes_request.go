package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttributesRequest Request Object
type ListAttributesRequest struct {

	// 自定义属性名称
	CustAttrName *string `json:"cust_attr_name,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// 自定义属性状态：0 未启用，1 已启用。
	Status *int32 `json:"status,omitempty"`
}

func (o ListAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListAttributesRequest", string(data)}, " ")
}
