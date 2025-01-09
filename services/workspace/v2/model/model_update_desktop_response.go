package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopResponse Response Object
type UpdateDesktopResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`
	HttpStatusCode              int     `json:"-"`
}

func (o UpdateDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesktopResponse", string(data)}, " ")
}
