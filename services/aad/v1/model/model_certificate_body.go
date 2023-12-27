package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateBody struct {

	// 操作类型。0 - 上传， 1 - 更换
	OpType int32 `json:"op_type"`

	// 证书名称
	CertName string `json:"cert_name"`

	// 证书文件。上传新证书(op_type为0)时必填，更换为已有证书(op_type为1)时置空
	CertFile *string `json:"cert_file,omitempty"`

	// 私钥文件。上传新证书(op_type为0)时必填，更换为已有证书(op_type为1)时置空
	CertKeyFile *string `json:"cert_key_file,omitempty"`

	// 域名id
	DomainId string `json:"domain_id"`
}

func (o CertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateBody struct{}"
	}

	return strings.Join([]string{"CertificateBody", string(data)}, " ")
}
