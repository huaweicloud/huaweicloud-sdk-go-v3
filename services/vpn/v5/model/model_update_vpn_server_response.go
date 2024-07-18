package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnServerResponse Response Object
type UpdateVpnServerResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnServerResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnServerResponse", string(data)}, " ")
}
