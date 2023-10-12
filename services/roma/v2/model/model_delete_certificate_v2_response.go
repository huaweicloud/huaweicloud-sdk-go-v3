package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateV2Response Response Object
type DeleteCertificateV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateV2Response struct{}"
	}

	return strings.Join([]string{"DeleteCertificateV2Response", string(data)}, " ")
}
