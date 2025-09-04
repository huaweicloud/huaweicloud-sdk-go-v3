package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceCertificateResponse Response Object
type DeleteDeviceCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceCertificateResponse", string(data)}, " ")
}
