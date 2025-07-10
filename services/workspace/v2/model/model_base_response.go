package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseResponse API响应基类。
type BaseResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`
}

func (o BaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseResponse struct{}"
	}

	return strings.Join([]string{"BaseResponse", string(data)}, " ")
}
