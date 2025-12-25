package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetryPolicyResponse Response Object
type CreateRetryPolicyResponse struct {

	// 状态码
	Code *string `json:"code,omitempty"`

	// 下发任务id
	Data *string `json:"data,omitempty"`

	// 状态信息
	Message *string `json:"message,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 请求状态
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateRetryPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyResponse", string(data)}, " ")
}
