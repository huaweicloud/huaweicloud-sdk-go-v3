package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeFunctionResponse Response Object
type InvokeFunctionResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// 函数执行结果
	Result *string `json:"result,omitempty"`

	// 函数执行返回日志
	Log *string `json:"log,omitempty"`

	// 函数执行返回状态
	Status *int32 `json:"status,omitempty"`

	XCffRequestId  *string `json:"X-Cff-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InvokeFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeFunctionResponse struct{}"
	}

	return strings.Join([]string{"InvokeFunctionResponse", string(data)}, " ")
}
