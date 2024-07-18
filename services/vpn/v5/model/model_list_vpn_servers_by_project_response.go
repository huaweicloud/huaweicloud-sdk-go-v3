package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnServersByProjectResponse Response Object
type ListVpnServersByProjectResponse struct {
	VpnServers *[]ShowServerResponse `json:"vpn_servers,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListVpnServersByProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnServersByProjectResponse struct{}"
	}

	return strings.Join([]string{"ListVpnServersByProjectResponse", string(data)}, " ")
}
