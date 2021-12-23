package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateCertificateV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateCertificateV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateCertificateV2Response", string(data)}, " ")
}
