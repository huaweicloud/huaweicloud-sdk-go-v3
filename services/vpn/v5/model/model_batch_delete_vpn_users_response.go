package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteVpnUsersResponse Response Object
type BatchDeleteVpnUsersResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o BatchDeleteVpnUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpnUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpnUsersResponse", string(data)}, " ")
}
