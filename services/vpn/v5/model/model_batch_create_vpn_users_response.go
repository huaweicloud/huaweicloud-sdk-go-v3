package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateVpnUsersResponse Response Object
type BatchCreateVpnUsersResponse struct {

	// 异常的用户列表
	InvalidUsers *[]InvalidVpnUser `json:"invalid_users,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o BatchCreateVpnUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnUsersResponse", string(data)}, " ")
}
