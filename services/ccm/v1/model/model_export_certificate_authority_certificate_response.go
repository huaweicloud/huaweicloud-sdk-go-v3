package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportCertificateAuthorityCertificateResponse struct {

	// 证书内容。  >  - 通过API请求本接口，证书内容中换行符已使用\"\\r\\n\"代替；  >  - 通过console端导出证书，将得到标准的PEM格式的证书文件。
	Certificate *string `json:"certificate,omitempty"`

	// 证书链内容，证书链中排列顺序（从上至下）：中间证书>...>根证书。 >  - 通过API请求本接口，证书链内容中换行符已使用\"\\r\\n\"代替； >  - 通过console端导出证书链，将得到标准的PEM格式的证书链文件。
	CertificateChain *string `json:"certificate_chain,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ExportCertificateAuthorityCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateAuthorityCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateAuthorityCertificateResponse", string(data)}, " ")
}
