package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePostpaidVgwSpecificationResponse Response Object
type UpdatePostpaidVgwSpecificationResponse struct {
	VpnGateway *UpdateResponseVpnGateway `json:"vpn_gateway,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdatePostpaidVgwSpecificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostpaidVgwSpecificationResponse struct{}"
	}

	return strings.Join([]string{"UpdatePostpaidVgwSpecificationResponse", string(data)}, " ")
}
