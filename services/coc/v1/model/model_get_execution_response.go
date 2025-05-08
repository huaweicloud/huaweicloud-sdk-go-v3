package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetExecutionResponse Response Object
type GetExecutionResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 返回数据
	Data *interface{} `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetExecutionResponse struct{}"
	}

	return strings.Join([]string{"GetExecutionResponse", string(data)}, " ")
}
