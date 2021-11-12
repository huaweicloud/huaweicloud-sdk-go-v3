package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStatisticsResponse struct {
	// 安全统计数据

	Body           *[]CountItem `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListStatisticsResponse", string(data)}, " ")
}
