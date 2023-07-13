package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushCertificateResponse Response Object
type PushCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PushCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushCertificateResponse struct{}"
	}

	return strings.Join([]string{"PushCertificateResponse", string(data)}, " ")
}
