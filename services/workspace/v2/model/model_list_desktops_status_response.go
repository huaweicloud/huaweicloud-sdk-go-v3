package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsStatusResponse Response Object
type ListDesktopsStatusResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`

	// 统计信息。
	Statics        *[]InstanceStatusStatistics `json:"statics,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListDesktopsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsStatusResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsStatusResponse", string(data)}, " ")
}
