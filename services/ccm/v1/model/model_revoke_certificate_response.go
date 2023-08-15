package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeCertificateResponse Response Object
type RevokeCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokeCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeCertificateResponse struct{}"
	}

	return strings.Join([]string{"RevokeCertificateResponse", string(data)}, " ")
}
