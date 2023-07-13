package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubMetricsResponse Response Object
type ListSubMetricsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 指标数据列表
	Result         *[]EventSubMetricsItem `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSubMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListSubMetricsResponse", string(data)}, " ")
}
