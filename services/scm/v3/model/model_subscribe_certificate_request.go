package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeCertificateRequest Request Object
type SubscribeCertificateRequest struct {
	Body *PurchaseCertificateRequestBody `json:"body,omitempty"`
}

func (o SubscribeCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeCertificateRequest struct{}"
	}

	return strings.Join([]string{"SubscribeCertificateRequest", string(data)}, " ")
}
