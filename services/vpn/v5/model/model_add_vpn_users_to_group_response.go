package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVpnUsersToGroupResponse Response Object
type AddVpnUsersToGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o AddVpnUsersToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpnUsersToGroupResponse struct{}"
	}

	return strings.Join([]string{"AddVpnUsersToGroupResponse", string(data)}, " ")
}
