package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitIpResponse Response Object
type CreateTransitIpResponse struct {
	TransitIp *TransitIp `json:"transit_ip,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitIpResponse", string(data)}, " ")
}
