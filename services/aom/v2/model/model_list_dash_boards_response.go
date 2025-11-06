package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashBoardsResponse Response Object
type ListDashBoardsResponse struct {

	// 当前页大小。
	PageSize *int32 `json:"page_size,omitempty"`

	// 仪表盘详情列表。
	Dashboards     *[]Dashboard `json:"dashboards,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDashBoardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashBoardsResponse struct{}"
	}

	return strings.Join([]string{"ListDashBoardsResponse", string(data)}, " ")
}
