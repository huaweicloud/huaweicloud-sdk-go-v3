package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageStatisticsResponse Response Object
type ListMessageStatisticsResponse struct {

	// 所有消息总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMessageStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListMessageStatisticsResponse", string(data)}, " ")
}
