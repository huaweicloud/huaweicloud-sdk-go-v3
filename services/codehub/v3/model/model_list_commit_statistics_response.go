package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCommitStatisticsResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *RepoCommitStatistics `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommitStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitStatisticsResponse", string(data)}, " ")
}
