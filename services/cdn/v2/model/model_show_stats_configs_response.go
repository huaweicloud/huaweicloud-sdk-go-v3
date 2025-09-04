package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatsConfigsResponse Response Object
type ShowStatsConfigsResponse struct {

	// 统计配置数量
	Total *int32 `json:"total,omitempty"`

	Data           *[]StatsConfigDetails `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowStatsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatsConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowStatsConfigsResponse", string(data)}, " ")
}
