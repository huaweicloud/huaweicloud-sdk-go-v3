package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingHistoryResponse Response Object
type ShowAutoScalingHistoryResponse struct {

	// 记录总数。
	TotalCount *string `json:"total_count,omitempty"`

	// 自动变配历史记录列表。
	Records        *[]AutoScalingRecordInfo `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAutoScalingHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingHistoryResponse", string(data)}, " ")
}
