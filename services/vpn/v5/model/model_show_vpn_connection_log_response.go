package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionLogResponse Response Object
type ShowVpnConnectionLogResponse struct {
	Logs *[]Log `json:"logs,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnConnectionLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionLogResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionLogResponse", string(data)}, " ")
}
