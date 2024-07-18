package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientCaResponse Response Object
type ShowClientCaResponse struct {
	ClientCaCertificate *QueryClientCaCertificateBody `json:"client_ca_certificate,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowClientCaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientCaResponse struct{}"
	}

	return strings.Join([]string{"ShowClientCaResponse", string(data)}, " ")
}
