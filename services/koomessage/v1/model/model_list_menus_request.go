package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMenusRequest struct {

	// 菜单ID。
	MenuId *string `json:"menu_id,omitempty"`

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 上线开始时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	OnlineBeginTime *string `json:"online_begin_time,omitempty"`

	// 上线结束时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	OnlineEndTime *string `json:"online_end_time,omitempty"`

	// 菜单状态。  - 1：未生效  - 2：已生效  - 3：已失效  - 4：服务号已冻结
	State *int32 `json:"state,omitempty"`

	// 一级菜单名称。
	MenuName *string `json:"menu_name,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMenusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMenusRequest struct{}"
	}

	return strings.Join([]string{"ListMenusRequest", string(data)}, " ")
}
