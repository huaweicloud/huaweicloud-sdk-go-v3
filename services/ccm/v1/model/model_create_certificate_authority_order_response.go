package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityOrderResponse Response Object
type CreateCertificateAuthorityOrderResponse struct {

	// 订单号。
	OrderId *string `json:"order_id,omitempty"`

	// 当前购买的CA证书ID列表。
	CaIds          *[]string `json:"ca_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateCertificateAuthorityOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityOrderResponse", string(data)}, " ")
}
