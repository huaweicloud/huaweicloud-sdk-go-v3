package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSpCertificateResponse Response Object
type DeleteSpCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSpCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSpCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteSpCertificateResponse", string(data)}, " ")
}
