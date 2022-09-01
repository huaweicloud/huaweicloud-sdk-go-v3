package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type InvokeFunctionResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 函数执行结果
	Result *string `json:"result,omitempty" xml:"result"`

	// 函数执行返回日志
	Log *string `json:"log,omitempty" xml:"log"`

	// 函数执行返回状态
	Status *int32 `json:"status,omitempty" xml:"status"`

	XCffRequestId  *string `json:"X-Cff-Request-Id,omitempty" xml:"X-Cff-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o InvokeFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeFunctionResponse struct{}"
	}

	return strings.Join([]string{"InvokeFunctionResponse", string(data)}, " ")
}
