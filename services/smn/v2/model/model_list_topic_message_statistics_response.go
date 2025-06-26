package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicMessageStatisticsResponse Response Object
type ListTopicMessageStatisticsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	Total *SumCountDetail `json:"total,omitempty"`

	// 周期内的发送详情列表
	Statistics     *[]StatisticsDetail `json:"statistics,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListTopicMessageStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicMessageStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicMessageStatisticsResponse", string(data)}, " ")
}
