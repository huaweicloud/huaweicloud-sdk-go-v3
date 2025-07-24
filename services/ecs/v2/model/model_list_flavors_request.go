package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// 可用区，需要指定可用区（AZ）的名称或者ID或者code。  可通过接口 [查询可用区列表接口](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ECS&api=NovaListAvailabilityZones) 获取，也可参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)获取。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 规格id
	FlavorId *string `json:"flavor_id,omitempty"`

	// 查询返回云服务器规格列表当前页面的数量。默认为1000
	Limit *int32 `json:"limit,omitempty"`

	// 从marker指定的flavor_id的下一条数据开始查询
	Marker *string `json:"marker,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
