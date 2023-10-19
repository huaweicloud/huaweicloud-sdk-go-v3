package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterRegionBandwidthsRequest Request Object
type ListInterRegionBandwidthsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据云连接的ID过滤列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据带宽包列表过滤域间带宽实例列表。
	BandwidthPackageId *[]string `json:"bandwidth_package_id,omitempty"`
}

func (o ListInterRegionBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterRegionBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListInterRegionBandwidthsRequest", string(data)}, " ")
}
