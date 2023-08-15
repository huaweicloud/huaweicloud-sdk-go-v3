package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitStatisticsResponse Response Object
type ListCommitStatisticsResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoCommitStatistics `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommitStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitStatisticsResponse", string(data)}, " ")
}
