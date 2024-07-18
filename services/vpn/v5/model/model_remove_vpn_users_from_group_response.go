package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveVpnUsersFromGroupResponse Response Object
type RemoveVpnUsersFromGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o RemoveVpnUsersFromGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpnUsersFromGroupResponse struct{}"
	}

	return strings.Join([]string{"RemoveVpnUsersFromGroupResponse", string(data)}, " ")
}
