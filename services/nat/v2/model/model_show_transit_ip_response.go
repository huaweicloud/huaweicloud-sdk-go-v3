package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTransitIpResponse struct {
	TransitIp *TransitIp `json:"transit_ip,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitIpResponse", string(data)}, " ")
}
