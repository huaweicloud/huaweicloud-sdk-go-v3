package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteNetworksRequest Request Object
type ListSiteNetworksRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 指定排序是升序还是降序(asc为升序，desc为降序)。
	SortDir *SortDir `json:"sort_dir,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据状态查询，可查询多个状态。
	State *[]SiteNetworkStateEnum `json:"state,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据GDW实例ID过滤列表。
	GlobalDcGatewayId *[]string `json:"global_dc_gateway_id,omitempty"`

	// 根据带宽包ID过滤。
	GlobalConnectionBandwidthId *[]string `json:"global_connection_bandwidth_id,omitempty"`

	// 分支连接的ID。
	ConnectionId *[]string `json:"connection_id,omitempty"`
}

func (o ListSiteNetworksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteNetworksRequest struct{}"
	}

	return strings.Join([]string{"ListSiteNetworksRequest", string(data)}, " ")
}
