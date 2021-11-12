package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCertificateRequestBody struct {
	// 证书名

	Name *string `json:"name,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
