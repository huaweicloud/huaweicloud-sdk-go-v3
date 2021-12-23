package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCertificateRequestBody struct {
	// 证书名称，证书名称只能由数字、字母、中划线、下划线和英文句点组成，长度不能超过64位字符

	Name *string `json:"name,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
