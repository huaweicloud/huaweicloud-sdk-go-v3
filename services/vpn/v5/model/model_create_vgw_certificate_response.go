package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVgwCertificateResponse Response Object
type CreateVgwCertificateResponse struct {
	Certificate *VpnGatewayCertificateConfig `json:"certificate,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVgwCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateVgwCertificateResponse", string(data)}, " ")
}
