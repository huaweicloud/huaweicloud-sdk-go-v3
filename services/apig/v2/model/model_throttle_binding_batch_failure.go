package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingBatchFailure struct {

	// 解除绑定失败的API和流控策略绑定关系ID
	BindId *string `json:"bind_id,omitempty"`

	// 解除绑定失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 解除绑定失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 解除绑定失败的API的ID
	ApiId *string `json:"api_id,omitempty"`

	// 解除绑定失败的API的名称
	ApiName *string `json:"api_name,omitempty"`
}

func (o ThrottleBindingBatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingBatchFailure struct{}"
	}

	return strings.Join([]string{"ThrottleBindingBatchFailure", string(data)}, " ")
}
