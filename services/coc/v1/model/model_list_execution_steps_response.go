package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionStepsResponse Response Object
type ListExecutionStepsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 返回数据
	Data *[]ExecutionStep `json:"data,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListExecutionStepsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionStepsResponse struct{}"
	}

	return strings.Join([]string{"ListExecutionStepsResponse", string(data)}, " ")
}
