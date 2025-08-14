package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationInstanceCertificateResponse Response Object
type DeleteApplicationInstanceCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationInstanceCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationInstanceCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationInstanceCertificateResponse", string(data)}, " ")
}
