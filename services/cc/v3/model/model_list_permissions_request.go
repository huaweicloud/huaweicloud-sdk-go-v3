package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionsRequest Request Object
type ListPermissionsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据描述查询，可查询多个描述。
	Description *[]string `json:"description,omitempty"`

	// 根据云连接的ID过滤列表。
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
