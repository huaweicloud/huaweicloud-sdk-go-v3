package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCertificateRequestBody struct {

	// 证书名称，证书名称只能由数字、字母、中划线、下划线和英文句点组成，长度不能超过256位字符
	Name string `json:"name"`

	// 证书文件，仅支持PEM格式的证书和私钥文件，且文件中的换行符应以\\n替换
	Content *string `json:"content,omitempty"`

	// 证书私钥，仅支持PEM格式的证书和私钥文件，且文件中的换行符应以\\n替换
	Key *string `json:"key,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
