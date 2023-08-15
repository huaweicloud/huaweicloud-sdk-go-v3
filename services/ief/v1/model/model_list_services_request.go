package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicesRequest Request Object
type ListServicesRequest struct {

	// 指定分页查询每页的行数，最大为100，默认值为10。
	Limit *int64 `json:"limit,omitempty"`

	// 指定要查询的偏移数量，默认为0。
	Offset *int64 `json:"offset,omitempty"`

	// 响应中查询到的服务将按照指定的字段进行排序
	Sorted *string `json:"sorted,omitempty"`

	// 服务名称
	Name *string `json:"name,omitempty"`

	// 按照相关的应用查询服务
	App *string `json:"app,omitempty"`

	// 铂金版实例ID
	IefInstanceId string `json:"ief-instance-id"`
}

func (o ListServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesRequest struct{}"
	}

	return strings.Join([]string{"ListServicesRequest", string(data)}, " ")
}
