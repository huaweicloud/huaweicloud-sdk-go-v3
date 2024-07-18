package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnServersByVgwResponse Response Object
type ListVpnServersByVgwResponse struct {
	VpnServers *[]ShowServerResponse `json:"vpn_servers,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListVpnServersByVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnServersByVgwResponse struct{}"
	}

	return strings.Join([]string{"ListVpnServersByVgwResponse", string(data)}, " ")
}
