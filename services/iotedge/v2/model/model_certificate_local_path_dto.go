package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CertificateLocalPathDto 证书文件的本地存储路径
type CertificateLocalPathDto struct {

	// 节点数字证书的本地存储路径。
	CertPath string `json:"cert_path"`

	// 证书私钥的本地存储路径。
	KeyPath string `json:"key_path"`
}

func (o CertificateLocalPathDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateLocalPathDto struct{}"
	}

	return strings.Join([]string{"CertificateLocalPathDto", string(data)}, " ")
}
