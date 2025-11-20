package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerCertificateResponse Response Object
type DeleteServerCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerCertificateResponse", string(data)}, " ")
}
