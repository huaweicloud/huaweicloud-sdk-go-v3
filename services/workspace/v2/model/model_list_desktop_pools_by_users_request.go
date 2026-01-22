package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolsByUsersRequest Request Object
type ListDesktopPoolsByUsersRequest struct {

	// 待查询的用户id列表
	UserIds *[]string `json:"user_ids,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-10，默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopPoolsByUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolsByUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolsByUsersRequest", string(data)}, " ")
}
