package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportCertificateAuthorityCsrResponse struct {

	// 证书签名请求内容，有以下两种情况：   - 通过API请求本接口，证书签名请求中换行符已使用\"\\r\\n\"代替；   - 通过console端导出证书签名请求，将得到标准的PEM格式的证书签名请求文件。
	Csr            *string `json:"csr,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCertificateAuthorityCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateAuthorityCsrResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateAuthorityCsrResponse", string(data)}, " ")
}
