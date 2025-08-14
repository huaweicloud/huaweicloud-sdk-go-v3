package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSpActiveCertificateResponse Response Object
type UpdateSpActiveCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSpActiveCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpActiveCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateSpActiveCertificateResponse", string(data)}, " ")
}
