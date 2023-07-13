package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPubMetricsResponse Response Object
type ListPubMetricsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 指标列表
	Result         *[]EventPubMetricsItem `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPubMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPubMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListPubMetricsResponse", string(data)}, " ")
}
