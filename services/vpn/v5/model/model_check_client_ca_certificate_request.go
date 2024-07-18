package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClientCaCertificateRequest Request Object
type CheckClientCaCertificateRequest struct {
	Body *ImportClientCaCertificateRequestBody `json:"body,omitempty"`
}

func (o CheckClientCaCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClientCaCertificateRequest struct{}"
	}

	return strings.Join([]string{"CheckClientCaCertificateRequest", string(data)}, " ")
}
