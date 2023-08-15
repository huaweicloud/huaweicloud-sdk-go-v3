package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterRegionBandwidthsRequest Request Object
type ListInterRegionBandwidthsRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤域间带宽实例列表。
	Id *[]string `json:"id,omitempty"`

	// 根据企业项目ID过滤域间带宽实例列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据云连接ID过滤域间带宽实例列表。
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
