package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCertificateResponse Response Object
type UpdateCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCertificateResponse", string(data)}, " ")
}
