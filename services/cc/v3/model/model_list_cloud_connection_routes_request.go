package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCloudConnectionRoutesRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 根据云连接路由ID过滤云连接路由条目列表。
	Id *string `json:"id,omitempty" xml:"id"`

	// 根据云连接实例ID过滤云连接路由条目列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty" xml:"cloud_connection_id"`

	// 根据网络实例ID过滤云连接路由条目列表。
	InstanceId *[]string `json:"instance_id,omitempty" xml:"instance_id"`

	// 根据Region ID过滤云连接路由条目列表。
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`
}

func (o ListCloudConnectionRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionRoutesRequest", string(data)}, " ")
}
