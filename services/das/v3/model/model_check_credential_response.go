package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCredentialResponse Response Object
type CheckCredentialResponse struct {

	// 测试结果
	CheckResult *bool `json:"check_result,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCredentialResponse struct{}"
	}

	return strings.Join([]string{"CheckCredentialResponse", string(data)}, " ")
}
