package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecyclingJobResponse Response Object
type ListRecyclingJobResponse struct {

	// 状态
	Success *bool `json:"success,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 错误码
	ErrCode *string `json:"err_code,omitempty"`

	Result *RecyclingJobsResult `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecyclingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecyclingJobResponse struct{}"
	}

	return strings.Join([]string{"ListRecyclingJobResponse", string(data)}, " ")
}
