package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingBatchFailure struct {

	// 解除绑定失败的API和流控策略绑定关系ID
	BindId *string `json:"bind_id,omitempty" xml:"bind_id"`

	// 解除绑定失败的错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 解除绑定失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 解除绑定失败的API的ID
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// 解除绑定失败的API的名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`
}

func (o ThrottleBindingBatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingBatchFailure struct{}"
	}

	return strings.Join([]string{"ThrottleBindingBatchFailure", string(data)}, " ")
}
