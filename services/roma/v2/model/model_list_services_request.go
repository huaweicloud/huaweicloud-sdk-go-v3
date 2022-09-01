package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServicesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 服务ID
	ServiceId *int32 `json:"service_id,omitempty" xml:"service_id"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty" xml:"service_name"`

	// 归属产品模板ID，product_template_id和product_id二选一
	ProductTemplateId *int32 `json:"product_template_id,omitempty" xml:"product_template_id"`

	// 归属产品ID，product_template_id和product_id二选一
	ProductId *int32 `json:"product_id,omitempty" xml:"product_id"`

	// 创建用户名
	CreatedUserName *string `json:"created_user_name,omitempty" xml:"created_user_name"`

	// 创建时间起始，格式timestamp(ms)，使用UTC时区
	CreatedDateStart *int64 `json:"created_date_start,omitempty" xml:"created_date_start"`

	// 创建时间截止，格式timestamp(ms)。使用UTC时区
	CreatedDateEnd *int64 `json:"created_date_end,omitempty" xml:"created_date_end"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesRequest struct{}"
	}

	return strings.Join([]string{"ListServicesRequest", string(data)}, " ")
}
