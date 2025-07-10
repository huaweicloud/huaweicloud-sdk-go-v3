package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolAuthorizedObjectsRequest Request Object
type ListDesktopPoolAuthorizedObjectsRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-100，默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopPoolAuthorizedObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolAuthorizedObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolAuthorizedObjectsRequest", string(data)}, " ")
}
