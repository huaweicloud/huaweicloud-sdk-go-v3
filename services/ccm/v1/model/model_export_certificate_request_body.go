package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportCertificateRequestBody struct {

	// 是否压缩。   - **true**   - **false**
	IsCompressed string `json:"is_compressed"`

	// 根据服务器类型选择下载证书的形式，支持以下五种类型：   - **APACHE** : apache服务器推荐使用此参数；   - **NGINX** : nginx服务器推荐使用此参数；   - **IIS** : windows服务器推荐使用此参数；   - **TOMCAT** : tomcat服务器推荐使用此参数；   - **OTHER** : 下载PEM格式证书，推荐使用此参数。
	Type string `json:"type"`

	// 是否国密GMT0009标准规范。当证书算法为SM2时传入才有效，若不传入，则默认为false。   - **true**   - **false**
	IsSmStandard *string `json:"is_sm_standard,omitempty"`

	// 设置用于加密私钥的密码。支持使用英文大小写字母、数字、特殊字符（例如,.+-_#）等。最大长度为32字节，若不传入，则默认不使用加密导出。
	Password *string `json:"password,omitempty"`
}

func (o ExportCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequestBody", string(data)}, " ")
}
