package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAutoLaunchGroupsRequest struct {

	// 查询返回智能购买组的数量限制 取值范围：1-1000。
	Limit *int32 `json:"limit,omitempty"`

	// 取值为上一页数据的最后一条记录的唯一标识
	Marker *string `json:"marker,omitempty"`

	// 智能购买组名称
	Name *string `json:"name,omitempty"`

	// 请求开始时间，按照时间范围过滤 按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ
	ValidSince *string `json:"valid_since,omitempty"`

	// 请求结束时间,按照时间范围过滤 按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ
	ValidUntil *string `json:"valid_until,omitempty"`
}

func (o ListAutoLaunchGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchGroupsRequest", string(data)}, " ")
}
