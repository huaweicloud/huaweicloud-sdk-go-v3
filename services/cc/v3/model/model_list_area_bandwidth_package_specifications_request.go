package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAreaBandwidthPackageSpecificationsRequest Request Object
type ListAreaBandwidthPackageSpecificationsRequest struct {

	// （索引位置，偏移量）， 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数（action为count时无此参数） 从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 根据本端大区ID过滤带宽包资源规格列表。
	LocalAreaId *[]string `json:"local_area_id,omitempty"`

	// 根据对端大区ID过滤带宽包资源规格列表。
	RemoteAreaId *[]string `json:"remote_area_id,omitempty"`
}

func (o ListAreaBandwidthPackageSpecificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreaBandwidthPackageSpecificationsRequest struct{}"
	}

	return strings.Join([]string{"ListAreaBandwidthPackageSpecificationsRequest", string(data)}, " ")
}
