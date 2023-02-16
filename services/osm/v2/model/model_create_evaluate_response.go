package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEvaluateResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 请求Id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEvaluateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvaluateResponse struct{}"
	}

	return strings.Join([]string{"CreateEvaluateResponse", string(data)}, " ")
}
