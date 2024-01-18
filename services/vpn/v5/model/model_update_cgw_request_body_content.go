package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCgwRequestBodyContent struct {

	// 网关名称
	Name *string `json:"name,omitempty"`

	CaCertificate *CaCertificateRequest `json:"ca_certificate,omitempty"`
}

func (o UpdateCgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateCgwRequestBodyContent", string(data)}, " ")
}
