package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPermissionsRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤授权列表。
	Id *[]string `json:"id,omitempty"`

	// 根据名称过滤授权列表。
	Name *[]string `json:"name,omitempty"`

	// 根据描述过滤授权列表。
	Description *[]string `json:"description,omitempty"`

	// 根据云连接实例ID过滤授权列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据实例ID过滤授权列表。
	InstanceId *[]string `json:"instance_id,omitempty"`
}

func (o ListPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionsRequest", string(data)}, " ")
}
