package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVgwCertificateResponse Response Object
type UpdateVgwCertificateResponse struct {
	Certificate *VpnGatewayCertificateConfig `json:"certificate,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVgwCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateVgwCertificateResponse", string(data)}, " ")
}
