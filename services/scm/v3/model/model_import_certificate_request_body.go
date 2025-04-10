package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportCertificateRequestBody struct {

	// 证书名称。字符长度为3~63位, 请输入英文字符，数字，下划线，中划线，英文句点。
	Name string `json:"name"`

	// 证书内容，可包含中间证书及根证书。若certificate_chain字段传入证书链，则该字段只取证书本身。回车换行需要使用转义字符\\n或者\\r\\n替换。
	Certificate string `json:"certificate"`

	// 证书链，非必填，可通过certificate字段传入。回车换行需要使用转义字符\\n或者\\r\\n替换。
	CertificateChain *string `json:"certificate_chain,omitempty"`

	// 证书私钥。 不能上传带有口令保护的私钥，回车换行需要使用转义字符\\n或者\\r\\n替换。
	PrivateKey string `json:"private_key"`

	// 是否允许上传相同证书。 - true：同意上传相同证书。 - false：不同意上传相同证书。
	DuplicateCheck *bool `json:"duplicate_check,omitempty"`

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件。 - 取值为“all” - 取值为“0” - 满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 可选参数，国密证书的加密证书内容。回车换行需要使用转义字符\\n或者\\r\\n替换。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// 可选参数，国密证书的加密私钥。 不能上传带有口令保护的私钥，回车换行需要使用转义字符\\n或者\\r\\n替换。
	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o ImportCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ImportCertificateRequestBody", string(data)}, " ")
}
