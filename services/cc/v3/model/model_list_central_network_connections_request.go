package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkConnectionsRequest Request Object
type ListCentralNetworkConnectionsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 指定排序是升序还是降序（asc为升序，desc为降序）。
	SortDir *SortDir `json:"sort_dir,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据状态查询，可查询多个状态。
	State *[]CentralNetworkConnectionStateEnum `json:"state,omitempty"`

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	// 根据带宽包ID过滤。
	GlobalConnectionBandwidthId *[]string `json:"global_connection_bandwidth_id,omitempty"`

	// 根据带宽类型查询。带宽类型包括： - BandwidthPackage (按带宽计费，需要绑定全域互联带宽，并指定分配带宽大小) - TestBandwidth (不收费的测试带宽，仅保留最小带宽，用于测试跨地域连通性）
	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type,omitempty"`

	// 连接类型，支持。
	ConnectionType *ConnectionTypeEnum `json:"connection_type,omitempty"`

	// 是否跨地域。
	IsCrossRegion *bool `json:"is_cross_region,omitempty"`
}

func (o ListCentralNetworkConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkConnectionsRequest", string(data)}, " ")
}
