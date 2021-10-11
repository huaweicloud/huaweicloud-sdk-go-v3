package model

import (
	"encoding/json"

	"strings"
)

type AclBindingBatchResultFailureResp struct {
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

func (o AclBindingBatchResultFailureResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AclBindingBatchResultFailureResp struct{}"
	}

	return strings.Join([]string{"AclBindingBatchResultFailureResp", string(data)}, " ")
}
