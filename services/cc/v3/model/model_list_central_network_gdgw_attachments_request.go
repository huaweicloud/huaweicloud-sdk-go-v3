package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkGdgwAttachmentsRequest Request Object
type ListCentralNetworkGdgwAttachmentsRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 指定排序是升序还是降序（asc为升序，desc为降序）。
	SortDir *SortDir `json:"sort_dir,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 根据名称查询，可查询多个名称。
	Name *[]string `json:"name,omitempty"`

	// 根据状态查询，可查询多个状态。
	State *[]CentralNetworkConnectionStateEnum `json:"state,omitempty"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 根据GDW实例ID过滤列表。
	GlobalDcGatewayId *[]string `json:"global_dc_gateway_id,omitempty"`
}

func (o ListCentralNetworkGdgwAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkGdgwAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkGdgwAttachmentsRequest", string(data)}, " ")
}
