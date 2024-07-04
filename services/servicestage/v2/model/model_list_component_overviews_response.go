package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentOverviewsResponse Response Object
type ListComponentOverviewsResponse struct {

	// 组件个数。
	Count *int32 `json:"count,omitempty"`

	// 组件部署信息列表。
	Components     *[]ComponentOverview `json:"components,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListComponentOverviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentOverviewsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentOverviewsResponse", string(data)}, " ")
}
