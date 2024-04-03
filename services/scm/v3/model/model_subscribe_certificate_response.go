package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeCertificateResponse Response Object
type SubscribeCertificateResponse struct {

	// 订单号。
	OrderId *string `json:"order_id,omitempty"`

	// 证书列表，详情请参见CertDetail字段数据结构说明。
	Cert           *[]CertDetail `json:"cert,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o SubscribeCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeCertificateResponse struct{}"
	}

	return strings.Join([]string{"SubscribeCertificateResponse", string(data)}, " ")
}
