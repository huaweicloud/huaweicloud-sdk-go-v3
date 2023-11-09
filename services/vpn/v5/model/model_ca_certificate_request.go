package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaCertificateRequest 对端网关CA证书
type CaCertificateRequest struct {

	// 对端网关CA证书内容
	Content *string `json:"content,omitempty"`
}

func (o CaCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaCertificateRequest struct{}"
	}

	return strings.Join([]string{"CaCertificateRequest", string(data)}, " ")
}
