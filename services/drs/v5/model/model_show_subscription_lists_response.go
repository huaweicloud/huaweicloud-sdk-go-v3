package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionListsResponse Response Object
type ShowSubscriptionListsResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount int32 `json:"total_count"`

	Jobs           *SubscriptionListResp `json:"jobs"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowSubscriptionListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionListsResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionListsResponse", string(data)}, " ")
}
