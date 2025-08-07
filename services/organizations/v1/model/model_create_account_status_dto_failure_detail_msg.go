package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountStatusDtoFailureDetailMsg 请求失败详细原因
type CreateAccountStatusDtoFailureDetailMsg struct {

	// 透传错误码
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 透传鉴权失败信息
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`
}

func (o CreateAccountStatusDtoFailureDetailMsg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountStatusDtoFailureDetailMsg struct{}"
	}

	return strings.Join([]string{"CreateAccountStatusDtoFailureDetailMsg", string(data)}, " ")
}
