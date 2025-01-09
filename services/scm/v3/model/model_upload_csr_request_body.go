package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadCsrRequestBody struct {

	// 自定义CSR名称。
	Name string `json:"name"`

	// 证书CSR文件。
	Csr string `json:"csr"`

	// 证书私钥文件。
	PrivateKey *string `json:"private_key,omitempty"`
}

func (o UploadCsrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCsrRequestBody struct{}"
	}

	return strings.Join([]string{"UploadCsrRequestBody", string(data)}, " ")
}
