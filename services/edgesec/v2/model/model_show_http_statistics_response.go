package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpStatisticsResponse Response Object
type ShowHttpStatisticsResponse struct {

	// 安全统计数据
	Body           *[]HttpStatisticsItem `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowHttpStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpStatisticsResponse", string(data)}, " ")
}
