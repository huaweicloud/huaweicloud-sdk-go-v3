package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublishedAppRequest Request Object
type ListPublishedAppRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 应用状态： * `NORMAL` - 正常状态。 * `FORBIDDEN` - 禁用状态。
	State *string `json:"state,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`
}

func (o ListPublishedAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublishedAppRequest struct{}"
	}

	return strings.Join([]string{"ListPublishedAppRequest", string(data)}, " ")
}
