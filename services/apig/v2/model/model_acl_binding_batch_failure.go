package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBindingBatchFailure struct {

	// 解除绑定失败的API和ACL绑定关系ID
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

func (o AclBindingBatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBindingBatchFailure struct{}"
	}

	return strings.Join([]string{"AclBindingBatchFailure", string(data)}, " ")
}
