package model

import (
	"encoding/json"

	"strings"
)

type ImportCertificateRequestBody struct {
	// 证书名称。字符长度为0~63位。

	Name string `json:"name"`
	// 证书内容。回车换行需要使用转义字符\\n或者\\r\\n替换。

	Certificate string `json:"certificate"`
	// 证书链。回车换行需要使用转义字符\\n或者\\r\\n替换。

	CertificateChain string `json:"certificate_chain"`
	// 证书私钥。 不能上传带有口令保护的私钥，回车换行需要使用转义字符\\n或者\\r\\n替换。

	PrivateKey string `json:"private_key"`
}

func (o ImportCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ImportCertificateRequestBody", string(data)}, " ")
}
