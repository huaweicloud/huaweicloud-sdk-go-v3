package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExternalIdPCertificateResponse Response Object
type DeleteExternalIdPCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExternalIdPCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExternalIdPCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteExternalIdPCertificateResponse", string(data)}, " ")
}
