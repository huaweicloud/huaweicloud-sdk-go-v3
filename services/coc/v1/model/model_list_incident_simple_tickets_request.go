package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentSimpleTicketsRequest Request Object
type ListIncidentSimpleTicketsRequest struct {

	// 列表查询接口分页参数，表示当前查询第几页，1表示查询第一页。
	Offset *int32 `json:"offset,omitempty"`

	// 列表查询接口分页参数，表示一页查询多少条数据，该值需大于等于1且为整数，表示分页查询时每一页最多查到的数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIncidentSimpleTicketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentSimpleTicketsRequest struct{}"
	}

	return strings.Join([]string{"ListIncidentSimpleTicketsRequest", string(data)}, " ")
}
