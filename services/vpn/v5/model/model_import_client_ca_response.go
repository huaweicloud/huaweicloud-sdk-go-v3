package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportClientCaResponse Response Object
type ImportClientCaResponse struct {
	ClientCaCertificate *ImportClientCaCertificateResponseBodyClientCaCertificate `json:"client_ca_certificate,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ImportClientCaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportClientCaResponse struct{}"
	}

	return strings.Join([]string{"ImportClientCaResponse", string(data)}, " ")
}
